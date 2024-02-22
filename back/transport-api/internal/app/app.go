package app

import (
	"transport-api/internal/adapters/grpcads"
	"transport-api/internal/adapters/grpccore"
)

type appImpl struct {
	core grpccore.CoreClient
	ads  grpcads.AdsClient
}

func NewApp(coreCli grpccore.CoreClient) App {
	return &appImpl{core: coreCli}
}
