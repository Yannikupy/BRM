package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

const (
	createEmployeeQuery = `
		INSERT INTO "employees" ("company_id", "first_name", "second_name", "email", "job_title", "department", "creation_date", "is_deleted") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING "id";`

	updateEmployeeQuery = `
		UPDATE "employees"
		SET "first_name" = $2,
		    "second_name" = $3,
		    "job_title" = $4,
		    "department" = $5
		WHERE "id" = $1 AND (NOT "is_deleted");`

	deleteEmployeeQuery = `
		UPDATE "employees"
		SET "is_deleted" = true
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getCompanyEmployeesQuery = `
		SELECT * FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "job_title" = $3)
			AND ((NOT $4) OR "department" = $5)
		LIMIT $6 OFFSET $7;`

	getEmployeeByNameQuery = `
		SELECT * FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted") AND ("first_name" LIKE $2 OR "second_name" LIKE $2)
		LIMIT $3 OFFSET $4;`

	getEmployeeByIdQuery = `
		SELECT * FROM "employees"
		WHERE "id" = $1 AND (NOT "is_deleted");`
)

func (c *coreRepoImpl) CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error) {
	// TODO добавить обработку ошибки unique email

	var employeeId uint
	if err := c.QueryRow(ctx, createEmployeeQuery,
		employee.CompanyId,
		employee.FirstName,
		employee.SecondName,
		employee.Email,
		employee.JobTitle,
		employee.Department,
		employee.CreationDate,
		employee.IsDeleted,
	).Scan(&employeeId); err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		employee.Id = employeeId
		return employee, nil
	}
}

func (c *coreRepoImpl) UpdateEmployee(ctx context.Context, employeeId uint, upd model.UpdateEmployee) (model.Employee, error) {
	if e, err := c.Exec(ctx, updateEmployeeQuery,
		employeeId,
		upd.FirstName,
		upd.SecondName,
		upd.JobTitle,
		upd.Department,
	); err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Employee{}, model.ErrEmployeeNotExists
	} else {
		return c.GetEmployeeById(ctx, employeeId)
	}
}

func (c *coreRepoImpl) DeleteEmployee(ctx context.Context, employeeId uint) error {
	if e, err := c.Exec(ctx, deleteEmployeeQuery,
		employeeId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrEmployeeNotExists
	} else {
		return nil
	}
}

func (c *coreRepoImpl) GetCompanyEmployees(ctx context.Context, companyId uint, filter model.FilterEmployee) ([]model.Employee, error) {
	rows, err := c.Query(ctx, getCompanyEmployeesQuery,
		companyId,
		filter.ByJobTitle,
		filter.JobTitle,
		filter.ByDepartment,
		filter.Department,
		filter.Limit,
		filter.Offset)
	if err != nil {
		return []model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		_ = rows.Scan(
			&e.Id,
			&e.CompanyId,
			&e.FirstName,
			&e.SecondName,
			&e.Email,
			&e.JobTitle,
			&e.Department,
			&e.CreationDate,
			&e.IsDeleted)
		employees = append(employees, e)
	}
	return employees, nil
}

func (c *coreRepoImpl) GetEmployeeByName(ctx context.Context, companyId uint, ebn model.EmployeeByName) ([]model.Employee, error) {
	rows, err := c.Query(ctx, getEmployeeByNameQuery,
		companyId,
		ebn.Pattern+"%",
		ebn.Limit,
		ebn.Offset)
	if err != nil {
		return []model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		_ = rows.Scan(
			&e.Id,
			&e.CompanyId,
			&e.FirstName,
			&e.SecondName,
			&e.Email,
			&e.JobTitle,
			&e.Department,
			&e.CreationDate,
			&e.IsDeleted)
		employees = append(employees, e)
	}
	return employees, nil
}

func (c *coreRepoImpl) GetEmployeeById(ctx context.Context, employeeId uint) (model.Employee, error) {
	row := c.QueryRow(ctx, getEmployeeByIdQuery, employeeId)
	var employee model.Employee
	if err := row.Scan(
		&employee.Id,
		&employee.CompanyId,
		&employee.FirstName,
		&employee.SecondName,
		&employee.Email,
		&employee.JobTitle,
		&employee.Department,
		&employee.CreationDate,
		&employee.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, model.ErrEmployeeNotExists
	} else if err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return employee, nil
	}
}
