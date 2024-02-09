package grpcserver

import (
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	company, err := s.App.GetCompany(ctx, uint(req.Id))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetCompanyResponse{
		Company: &pb.Company{
			Id:           uint64(company.Id),
			Name:         company.Name,
			Description:  company.Description,
			Industry:     uint64(company.Industry),
			OwnerId:      uint64(company.OwnerId),
			Rating:       company.Rating,
			CreationDate: company.CreationDate,
			IsDeleted:    company.IsDeleted,
		},
	}, nil
}

func (s *Server) CreateCompanyAndOwner(ctx context.Context, req *pb.CreateCompanyAndOwnerRequest) (*pb.CreateCompanyAndOwnerResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*empty.Empty, error) {
	// TODO implement
	return &empty.Empty{}, nil
}
