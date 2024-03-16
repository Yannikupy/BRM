package app

import (
	"transport-api/internal/adapters/grpcads"
	"transport-api/internal/adapters/grpccore"
	"transport-api/internal/adapters/grpcleads"
	"transport-api/internal/adapters/grpcstats"
)

type appImpl struct {
	core  grpccore.CoreClient
	ads   grpcads.AdsClient
	leads grpcleads.LeadsClient
	stats grpcstats.StatsClient
}

func NewApp(
	coreCli grpccore.CoreClient,
	adsCli grpcads.AdsClient,
	leadsCli grpcleads.LeadsClient,
	statsCli grpcstats.StatsClient,
) App {
	return &appImpl{
		core:  coreCli,
		ads:   adsCli,
		leads: leadsCli,
		stats: statsCli,
	}
}
