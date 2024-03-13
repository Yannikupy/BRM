package core_repo

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type CoreRepo interface {
	GetCompanyAbsoluteRating(ctx context.Context, companyId uint64) (float64, error)
	GetCompanyRelativeRating(ctx context.Context, companyId uint64) (float64, error)
}

func New(conn *pgx.Conn) CoreRepo {
	return &coreRepoImpl{
		Conn: *conn,
	}
}
