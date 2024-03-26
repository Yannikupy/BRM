package core_repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CoreRepo interface {
	GetCompanyAbsoluteRating(ctx context.Context, companyId uint64) (float64, error)
	GetCompanyRelativeRating(ctx context.Context, companyId uint64) (float64, error)
}

func New(pool *pgxpool.Pool) CoreRepo {
	return &coreRepoImpl{
		Pool: pool,
	}
}
