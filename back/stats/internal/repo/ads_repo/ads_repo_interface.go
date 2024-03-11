package ads_repo

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type AdsRepo interface {
	GetActiveAdsAmount(ctx context.Context, companyId uint64) (uint, error)
}

func New(conn *pgx.Conn) AdsRepo {
	return &adsRepoImpl{
		Conn: *conn,
	}
}
