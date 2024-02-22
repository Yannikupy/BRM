package grpcads

import (
	"context"
	"transport-api/internal/model/ads"
)

type AdsClient interface {
	GetAdById(ctx context.Context, id uint64) (ads.Ad, error)
	GetAdsList(ctx context.Context, params ads.ListParams) ([]ads.Ad, error)
	CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad ads.Ad) (ads.Ad, error)
	UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd ads.UpdateAd) (ads.Ad, error)
	DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error

	CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (ads.Response, error)
	GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]ads.Response, error)
}
