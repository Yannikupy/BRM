package app

import (
	"brm-leads/internal/adapters/grpcads"
	"brm-leads/internal/adapters/grpccore"
	"brm-leads/internal/model"
	"brm-leads/internal/repo"
	"brm-leads/pkg/logger"
	"context"
)

type App interface {
	CreateLead(ctx context.Context, adId uint64, clientCompany uint64, clientEmployee uint64) (model.Lead, error)
	GetLeads(ctx context.Context, companyId uint64, employeeId uint64, filter model.Filter) ([]model.Lead, error)
	GetLeadById(ctx context.Context, companyId uint64, employeeId uint64, leadId uint64) (model.Lead, error)
	UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd model.UpdateLead) (model.Lead, error)
	DeleteLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64) error

	GetStatuses(ctx context.Context) (map[string]uint64, error)
	GetStatusById(ctx context.Context, id uint64) (string, error)
}

func New(repo repo.LeadsRepo, core grpccore.CoreClient, ads grpcads.AdsClient, logs logger.Logger) App {
	return &appImpl{
		leadsRepo:            repo,
		core:                 core,
		ads:                  ads,
		newLeadDefaultStatus: 1,
		logs:                 logs,
	}
}
