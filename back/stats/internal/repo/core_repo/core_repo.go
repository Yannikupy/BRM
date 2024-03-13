package core_repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"stats/internal/model"
)

type coreRepoImpl struct {
	pgx.Conn
}

const (
	getCompanyAbsoluteRatingQuery = `
		SELECT "rating" FROM "companies"
		WHERE "id" = $1 AND NOT "is_deleted";`

	getCompaniesAmountQuery = `
		SELECT COUNT(*) FROM "companies";`

	getComapniesWithLessRating = `
		SELECT COUNT(*) FROM "companies"
		WHERE "rating" <= $1 AND NOT "is_deleted";`
)

func (c *coreRepoImpl) GetCompanyAbsoluteRating(ctx context.Context, companyId uint64) (float64, error) {
	row := c.QueryRow(ctx, getCompanyAbsoluteRatingQuery, companyId)
	var rating float64
	if err := row.Scan(&rating); errors.Is(err, pgx.ErrNoRows) {
		return .0, model.ErrCompanyNotExists
	} else if err != nil {
		return .0, errors.Join(model.ErrCoreDatabase, err)
	} else {
		return rating, nil
	}
}

func (c *coreRepoImpl) GetCompanyRelativeRating(ctx context.Context, companyId uint64) (float64, error) {
	row := c.QueryRow(ctx, getCompaniesAmountQuery)
	var amount float64
	if err := row.Scan(&amount); errors.Is(err, pgx.ErrNoRows) {
		return .0, model.ErrCompanyNotExists
	} else if err != nil {
		return .0, errors.Join(model.ErrCoreDatabase, err)
	}

	rating, err := c.GetCompanyAbsoluteRating(ctx, companyId)
	if err != nil {
		return .0, err
	}

	row = c.QueryRow(ctx, getComapniesWithLessRating, rating)
	var lessRatingsCompaniesAmount float64
	if err = row.Scan(&lessRatingsCompaniesAmount); err != nil {
		return .0, errors.Join(model.ErrCoreDatabase, err)
	}

	return lessRatingsCompaniesAmount / amount, nil
}
