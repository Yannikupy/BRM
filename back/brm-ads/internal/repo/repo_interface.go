package repo

import (
	"brm-ads/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type AdRepo interface {
	GetAdById(ctx context.Context, id uint64) (model.Ad, error)
	GetAdsList(ctx context.Context, params model.AdsListParams) ([]model.Ad, error)
	CreateAd(ctx context.Context, ad model.Ad) (model.Ad, error)
	UpdateAd(ctx context.Context, adId uint64, upd model.UpdateAd) (model.Ad, error)
	DeleteAd(ctx context.Context, adId uint64) error

	CreateResponse(ctx context.Context, resp model.Response) (model.Response, error)
	GetResponses(ctx context.Context, companyId uint64, limit uint, offset uint) ([]model.Response, error)

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}

func New(conn *pgx.Conn) AdRepo {
	return &adRepoImpl{
		Conn: *conn,
	}
}
