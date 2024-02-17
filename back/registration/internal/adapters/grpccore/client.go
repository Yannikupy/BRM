package grpccore

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"registration/internal/adapters/grpccore/pb"
	"registration/internal/model"
)

type coreClientImpl struct {
	cli pb.CoreServiceClient
}

func NewCoreClient(ctx context.Context, addr string) (CoreClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &coreClientImpl{}, fmt.Errorf("grpc core client: %w", err)
	} else {
		return &coreClientImpl{
			cli: pb.NewCoreServiceClient(conn),
		}, nil
	}
}

func (c *coreClientImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	resp, err := c.cli.CreateCompanyAndOwner(ctx, &pb.CreateCompanyAndOwnerRequest{
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
		Owner: &pb.Employee{
			Id:           uint64(owner.Id),
			CompanyId:    uint64(owner.CompanyId),
			FirstName:    owner.FirstName,
			SecondName:   owner.SecondName,
			Email:        owner.Email,
			Password:     owner.Password,
			JobTitle:     owner.JobTitle,
			Department:   owner.Department,
			CreationDate: owner.CreationDate,
			IsDeleted:    owner.IsDeleted,
		},
	})

	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return model.Company{}, model.Employee{}, model.ErrCoreError
		default:
			return model.Company{}, model.Employee{}, model.ErrCoreUnknown
		}
	}

	return model.Company{
			Id:           uint(resp.Company.Id),
			Name:         resp.Company.Name,
			Description:  resp.Company.Description,
			Industry:     uint(resp.Company.Industry),
			OwnerId:      uint(resp.Company.OwnerId),
			Rating:       resp.Company.Rating,
			CreationDate: resp.Company.CreationDate,
			IsDeleted:    resp.Company.IsDeleted,
		},
		model.Employee{
			Id:           uint(resp.Owner.Id),
			CompanyId:    uint(resp.Owner.CompanyId),
			FirstName:    resp.Owner.FirstName,
			SecondName:   resp.Owner.SecondName,
			Email:        resp.Owner.Email,
			Password:     resp.Owner.Password,
			JobTitle:     resp.Owner.JobTitle,
			Department:   resp.Owner.Department,
			CreationDate: resp.Owner.CreationDate,
			IsDeleted:    resp.Owner.IsDeleted,
		}, nil
}
