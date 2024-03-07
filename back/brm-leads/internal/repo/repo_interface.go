package repo

import (
	"brm-leads/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type LeadsRepo interface {
	GetLeadById(ctx context.Context, id uint64) (model.Lead, error)
	GetLeads(ctx context.Context, companyId uint64, filter model.Filter) ([]model.Lead, error)
	CreateLead(ctx context.Context, lead model.Lead) (model.Lead, error)
	UpdateLead(ctx context.Context, id uint64, upd model.UpdateLead) (model.Lead, error)

	GetStatuses(ctx context.Context) (map[string]uint64, error)
}

func New(conn *pgx.Conn) LeadsRepo {
	return &leadRepoImpl{
		Conn: *conn,
	}
}
