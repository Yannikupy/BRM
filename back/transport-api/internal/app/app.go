package app

import (
	"transport-api/internal/adapters/grpccore"
)

type appImpl struct {
	core grpccore.CoreClient
}

func NewApp(coreCli grpccore.CoreClient) App {
	return &appImpl{core: coreCli}
}
