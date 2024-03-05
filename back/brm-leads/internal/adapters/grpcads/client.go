package grpcads

import (
	"brm-leads/internal/adapters/grpcads/pb"
	"brm-leads/internal/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type adsClientImpl struct {
	cli pb.AdsServiceClient
}

func (a *adsClientImpl) GetAdData(ctx context.Context, adId uint64) (model.AdData, error) {
	resp, err := a.cli.GetAdById(ctx, &pb.GetAdByIdRequest{Id: adId})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.AdData{}, model.ErrAdNotExists
		case codes.ResourceExhausted:
			return model.AdData{}, model.ErrAdsError
		default:
			return model.AdData{}, model.ErrAdsError
		}
	}
	return model.AdData{
		Price:       uint(resp.Ad.Price),
		Responsible: resp.Ad.Responsible,
		CompanyId:   resp.Ad.CompanyId,
	}, nil
}
