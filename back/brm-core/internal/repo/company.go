package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"strconv"
)

const (
	getCompanyQuery = `
		SELECT * FROM "companies"
		WHERE "id" = $1 AND (NOT "is_deleted");`

	createCompanyQuery = `
		INSERT INTO "companies" ("name", "description", "industry", "owner_id", "rating", "creation_date", "is_deleted") 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING "id";`

	createOwnerQuery = `
		INSERT INTO "employees" ("company_id", "first_name", "second_name", "email", "job_title", "department", "creation_date", "is_deleted")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING "id";`

	setCompanyOwnerIdQuery = `
		UPDATE "companies"
		SET "owner_id" = $2
		WHERE "id" = $1;`

	updateCompanyQuery = `
		UPDATE "companies"
		SET "name" = $2,
		    "description" = $3,
		    "industry" = $4,
		    "owner_id" = $5
		WHERE "id" = $1 AND (NOT "is_deleted");`

	deleteCompanyQuery = `
		UPDATE "companies"
		SET "is_deleted" = true
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getIndustriesListQuery = `
		SELECT * FROM "industries";`

	getIndustryByIdQuery = `
		SELECT "name" FROM "industries"
		WHERE "id" = $1;`
)

func (c *coreRepoImpl) GetCompany(ctx context.Context, id uint64) (model.Company, error) {
	row := c.QueryRow(ctx, getCompanyQuery, id)
	var company model.Company
	if err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.Industry,
		&company.OwnerId,
		&company.Rating,
		&company.CreationDate,
		&company.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Company{}, model.ErrCompanyNotExists
	} else if err != nil {
		return model.Company{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return company, nil
	}
}

func (c *coreRepoImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	var companyId, ownerId uint64
	var pgErr *pgconn.PgError
	if err := c.QueryRow(ctx, createCompanyQuery,
		company.Name,
		company.Description,
		company.Industry,
		0,
		company.Rating,
		company.CreationDate,
		company.IsDeleted,
	).Scan(&companyId); err != nil {
		return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}

	company.Id = companyId
	owner.CompanyId = companyId

	if err := c.QueryRow(ctx, createOwnerQuery,
		owner.CompanyId,
		owner.FirstName,
		owner.SecondName,
		owner.Email,
		owner.JobTitle,
		owner.Department,
		owner.CreationDate,
		owner.IsDeleted,
	).Scan(&ownerId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate primary key error
			return model.Company{}, model.Employee{}, model.ErrEmailRegistered
		default:
			return model.Company{}, model.Employee{}, model.ErrServiceError
		}
	} else if err != nil {
		return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}

	owner.Id = ownerId
	company.OwnerId = ownerId

	if _, err := c.Exec(ctx, setCompanyOwnerIdQuery, companyId, ownerId); err != nil {
		return model.Company{}, model.Employee{}, model.ErrDatabaseError
	}

	return company, owner, nil
}

func (c *coreRepoImpl) UpdateCompany(ctx context.Context, companyId uint64, upd model.UpdateCompany) (model.Company, error) {
	if e, err := c.Exec(ctx, updateCompanyQuery,
		companyId,
		upd.Name,
		upd.Description,
		upd.Industry,
		upd.OwnerId,
	); err != nil {
		return model.Company{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Company{}, model.ErrCompanyNotExists
	} else {
		return c.GetCompany(ctx, companyId)
	}
}

func (c *coreRepoImpl) DeleteCompany(ctx context.Context, companyId uint64) error {
	if e, err := c.Exec(ctx, deleteCompanyQuery,
		companyId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrCompanyNotExists
	} else {
		return nil
	}
}

func (c *coreRepoImpl) GetIndustriesList(ctx context.Context) (map[string]string, error) {
	rows, err := c.Query(ctx, getIndustriesListQuery)
	if err != nil {
		return map[string]string{}, model.ErrDatabaseError
	}
	defer rows.Close()

	industries := make(map[string]string)
	for rows.Next() {
		var id int
		var industry string
		_ = rows.Scan(&id, &industry)
		industries[strconv.Itoa(id)] = industry
	}
	return industries, nil
}

func (c *coreRepoImpl) GetIndustryById(ctx context.Context, id uint64) (string, error) {
	row := c.QueryRow(ctx, getIndustryByIdQuery, id)
	var industry string
	if err := row.Scan(&industry); errors.Is(err, pgx.ErrNoRows) {
		return "", model.ErrIndustryNotExists
	} else if err != nil {
		return "", errors.Join(model.ErrDatabaseError, err)
	} else {
		return industry, nil
	}
}
