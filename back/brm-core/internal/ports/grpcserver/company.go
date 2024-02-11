package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

func companyToModelCompany(company *pb.Company) model.Company {
	if company == nil {
		return model.Company{}
	}
	return model.Company{
		Id:           uint(company.Id),
		Name:         company.Name,
		Description:  company.Description,
		Industry:     uint(company.Industry),
		OwnerId:      uint(company.OwnerId),
		Rating:       company.Rating,
		CreationDate: time.Unix(company.CreationDate, 0),
		IsDeleted:    company.IsDeleted,
	}
}

func modelCompanyToCompany(company model.Company) *pb.Company {
	if company.Id == 0 {
		return nil
	}
	return &pb.Company{
		Id:           uint64(company.Id),
		Name:         company.Name,
		Description:  company.Description,
		Industry:     uint64(company.Industry),
		OwnerId:      uint64(company.OwnerId),
		Rating:       company.Rating,
		CreationDate: company.CreationDate.UTC().Unix(),
		IsDeleted:    company.IsDeleted,
	}
}

func (s *Server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	company, err := s.App.GetCompany(ctx, uint(req.Id))
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetCompanyResponse{
		Company: modelCompanyToCompany(company),
	}, nil
}

func (s *Server) CreateCompanyAndOwner(ctx context.Context, req *pb.CreateCompanyAndOwnerRequest) (*pb.CreateCompanyAndOwnerResponse, error) {
	company, owner, err := s.App.CreateCompanyAndOwner(ctx,
		companyToModelCompany(req.Company),
		employeeToModelEmployee(req.Owner),
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateCompanyAndOwnerResponse{
		Company: modelCompanyToCompany(company),
		Owner:   modelEmployeeToEmployee(owner),
	}, nil
}

func (s *Server) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	company, err := s.App.UpdateCompany(ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		model.UpdateCompany{
			Name:        req.Upd.Name,
			Description: req.Upd.Description,
			Industry:    uint(req.Upd.Industry),
			OwnerId:     uint(req.Upd.OwnerId),
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateCompanyResponse{
		Company: modelCompanyToCompany(company),
	}, nil
}

func (s *Server) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*empty.Empty, error) {
	if err := s.App.DeleteCompany(ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
	); err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}
