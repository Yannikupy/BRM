package leads_repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"stats/internal/model"
)

type LeadsRepo interface {
	GetMainPageLeadsStats(ctx context.Context, companyId uint64) (model.MainPageStats, error)
}

func New(conn *pgx.Conn) LeadsRepo {
	return &leadsRepoImpl{
		Conn: *conn,
	}
}
