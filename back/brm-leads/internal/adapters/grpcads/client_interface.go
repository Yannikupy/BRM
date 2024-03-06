package grpcads

import (
	"brm-leads/internal/model"
	"context"
)

type AdsClient interface {
	GetAdData(ctx context.Context, adId uint64) (model.AdData, error)
}
