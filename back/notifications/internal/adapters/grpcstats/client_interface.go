package grpcstats

import "context"

type StatsClient interface {
	SubmitClosedLead(ctx context.Context, producerCompanyId uint64, submit bool) error
}
