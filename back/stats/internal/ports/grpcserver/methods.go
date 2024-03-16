package grpcserver

import (
	"context"
	"stats/internal/ports/grpcserver/pb"
)

func (s *Server) GetCompanyMainPage(ctx context.Context, req *pb.GetCompanyMainPageRequest) (*pb.GetCompanyMainPageResponse, error) {
	resp, err := s.a.GetCompanyMainPageStats(ctx, req.CompanyId)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetCompanyMainPageResponse{Data: &pb.Data{
		ActiveLeadsAmount:     uint64(resp.ActiveLeadsAmount),
		ActiveLeadsPrice:      uint64(resp.ActiveLeadsPrice),
		ClosedLeadsAmount:     uint64(resp.ClosedLeadsAmount),
		ClosedLeadsPrice:      uint64(resp.ClosedLeadsPrice),
		ActiveAdsAmount:       uint64(resp.ActiveAdsAmount),
		CompanyAbsoluteRating: resp.CompanyAbsoluteRating,
		CompanyRelativeRating: resp.CompanyRelativeRating,
	}}, nil
}
