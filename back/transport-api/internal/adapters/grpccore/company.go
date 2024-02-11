package grpccore

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToCompany(company *pb.Company) core.Company {
	if company == nil {
		return core.Company{}
	}
	return core.Company{
		Id:           uint(company.Id),
		Name:         company.Name,
		Description:  company.Description,
		Industry:     uint(company.Industry),
		OwnerId:      uint(company.OwnerId),
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}
}

func (c *coreClientImpl) GetCompany(ctx context.Context, id uint) (core.Company, error) {
	resp, err := c.cli.GetCompany(ctx, &pb.GetCompanyRequest{Id: uint64(id)})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Company{}, model.ErrCompanyNotExists
		case codes.ResourceExhausted:
			return core.Company{}, model.ErrCoreError
		default:
			return core.Company{}, model.ErrCoreUnknown
		}
	}
	return respToCompany(resp.Company), nil
}

func (c *coreClientImpl) CreateCompanyAndOwner(ctx context.Context, company core.Company, owner core.Employee) (core.Company, core.Employee, error) {
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
		Owner: employeeToRequest(owner),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return core.Company{}, core.Employee{}, model.ErrCoreError
		default:
			return core.Company{}, core.Employee{}, model.ErrCoreUnknown
		}
	}

	return respToCompany(resp.Company), respToEmployee(resp.Owner), nil
}

func (c *coreClientImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd core.UpdateCompany) (core.Company, error) {
	resp, err := c.cli.UpdateCompany(ctx, &pb.UpdateCompanyRequest{
		CompanyId: uint64(companyId),
		OwnerId:   uint64(ownerId),
		Upd: &pb.UpdateCompanyFields{
			Name:        upd.Name,
			Description: upd.Description,
			Industry:    int64(upd.Industry),
			OwnerId:     uint64(upd.OwnerId),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Company{}, model.ErrCompanyNotExists
		case codes.PermissionDenied:
			return core.Company{}, model.ErrAuthorization
		case codes.ResourceExhausted:
			return core.Company{}, model.ErrCoreError
		default:
			return core.Company{}, model.ErrCoreUnknown
		}
	}
	return respToCompany(resp.Company), nil
}

func (c *coreClientImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error {
	_, err := c.cli.DeleteCompany(ctx, &pb.DeleteCompanyRequest{
		CompanyId: uint64(companyId),
		OwnerId:   uint64(ownerId),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrCompanyNotExists
		case codes.PermissionDenied:
			return model.ErrAuthorization
		case codes.ResourceExhausted:
			return model.ErrCoreError
		default:
			return model.ErrCoreUnknown
		}
	}
	return nil
}
