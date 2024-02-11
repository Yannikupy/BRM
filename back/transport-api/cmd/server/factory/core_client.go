package factory

import (
	"context"
	"transport-api/internal/adapters/grpccore"
)

func CreateCoreClient(ctx context.Context) grpccore.CoreClient {
	return nil
}
