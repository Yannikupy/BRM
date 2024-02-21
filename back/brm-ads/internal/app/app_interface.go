package app

import (
	"brm-ads/internal/model"
	"context"
)

type App interface {
	GetAdById(ctx context.Context, id uint64) (model.Ad, error)
	GetAdsList(ctx context.Context, params model.AdsListParams) ([]model.Ad, error)
	CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad model.Ad) (model.Ad, error)
	UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd model.UpdateAd) (model.Ad, error)
	DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error

	CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (model.Response, error)
	GetResponses(ctx context.Context, companyId uint64, employeeId uint64) ([]model.Response, error)
}
