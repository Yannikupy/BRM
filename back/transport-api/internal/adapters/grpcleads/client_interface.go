package grpcleads

import (
	"context"
	"transport-api/internal/model/leads"
)

type LeadsClient interface {
	GetLeads(ctx context.Context, companyId uint64, employeeId uint64, filter leads.Filter) ([]leads.Lead, error)
	GetLeadById(ctx context.Context, companyId uint64, employeeId uint64, leadId uint64) (leads.Lead, error)
	UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd leads.UpdateLead) (leads.Lead, error)

	GetStatuses(ctx context.Context) (map[string]uint64, error)
	GetStatusById(ctx context.Context, id uint64) (string, error)
}
