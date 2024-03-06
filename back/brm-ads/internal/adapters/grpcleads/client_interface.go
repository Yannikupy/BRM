package grpcleads

import "context"

type LeadsClient interface {
	CreateLead(ctx context.Context, adId uint64, clientCompany uint64, clientEmployee uint64) error
}
