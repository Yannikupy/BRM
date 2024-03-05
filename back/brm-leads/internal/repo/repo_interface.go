package repo

import (
	"brm-leads/internal/model"
	"context"
)

type LeadsRepo interface {
	GetLeadById(ctx context.Context, companyId uint64, id uint64) (model.Lead, error)
	GetLeads(ctx context.Context, companyId uint64, filter model.Filter) ([]model.Lead, error)
	CreateLead(ctx context.Context, lead model.Lead) (model.Lead, error)
	UpdateLead(ctx context.Context, companyId uint64, id uint64, upd model.UpdateLead) (model.Lead, error)

	GetStatuses(ctx context.Context) (map[string]uint64, error)
	GetStatusById(ctx context.Context, id uint64) (string, error)
}
