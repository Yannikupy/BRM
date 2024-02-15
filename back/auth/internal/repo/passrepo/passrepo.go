package passrepo

import (
	"auth/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type passRepoImpl struct {
	pgx.Conn
}

const (
	createEmployeeQuery = `
		INSERT INTO "passwords" ("email", "password", "employee_id", "company_id")
		VALUES ($1, $2, $3, $4);`

	getEmployeeQuery = `
		SELECT * FROM "passwords"
		WHERE "email" = $1;`

	deleteEmployeeQuery = `
		DELETE FROM "passwords"
		WHERE "email" = $1;`
)

func (p *passRepoImpl) CreateEmployee(ctx context.Context, employee model.Employee) error {
	if _, err := p.Exec(ctx, createEmployeeQuery,
		employee.Email,
		employee.Password,
		employee.EmployeeId,
		employee.CompanyId,
	); err != nil {
		return errors.Join(model.ErrPassRepoError, err)
	} else {
		return nil
	}
}

func (p *passRepoImpl) GetEmployee(ctx context.Context, email string) (model.Employee, error) {
	row := p.QueryRow(ctx, getEmployeeQuery, email)
	var employee model.Employee
	if err := row.Scan(
		&employee.Email,
		&employee.Password,
		&employee.EmployeeId,
		&employee.CompanyId,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, model.ErrEmployeeNotExists
	} else if err != nil {
		return model.Employee{}, errors.Join(model.ErrPassRepoError, err)
	} else {
		return employee, nil
	}
}

func (p *passRepoImpl) DeleteEmployee(ctx context.Context, email string) error {
	if e, err := p.Exec(ctx, deleteEmployeeQuery, email); err != nil {
		return errors.Join(model.ErrPassRepoError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrEmployeeNotExists
	} else {
		return nil
	}
}
