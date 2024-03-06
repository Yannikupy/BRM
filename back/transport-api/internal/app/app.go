package app

import (
	"transport-api/internal/adapters/grpcads"
	"transport-api/internal/adapters/grpccore"
	"transport-api/internal/adapters/grpcleads"
)

type appImpl struct {
	core  grpccore.CoreClient
	ads   grpcads.AdsClient
	leads grpcleads.LeadsClient
}

func NewApp(coreCli grpccore.CoreClient, adsCli grpcads.AdsClient, leadsCli grpcleads.LeadsClient) App {
	return &appImpl{
		core:  coreCli,
		ads:   adsCli,
		leads: leadsCli,
	}
}
