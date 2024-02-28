package grpccore

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToCompany(company *pb.Company) core.Company {
	if company == nil {
		return core.Company{}
	}
	return core.Company{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}
}

func (c *coreClientImpl) GetCompany(ctx context.Context, id uint64) (core.Company, error) {
	resp, err := c.cli.GetCompany(ctx, &pb.GetCompanyRequest{Id: id})
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

func (c *coreClientImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error) {
	resp, err := c.cli.UpdateCompany(ctx, &pb.UpdateCompanyRequest{
		CompanyId: companyId,
		OwnerId:   ownerId,
		Upd: &pb.UpdateCompanyFields{
			Name:        upd.Name,
			Description: upd.Description,
			Industry:    int64(upd.Industry),
			OwnerId:     upd.OwnerId,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			// костыль, ну а чё поделать
			if strings.Contains(err.Error(), "company") {
				return core.Company{}, model.ErrCompanyNotExists
			} else {
				return core.Company{}, model.ErrIndustryNotExists
			}
		case codes.PermissionDenied:
			return core.Company{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Company{}, model.ErrCoreError
		default:
			return core.Company{}, model.ErrCoreUnknown
		}
	}
	return respToCompany(resp.Company), nil
}

func (c *coreClientImpl) DeleteCompany(ctx context.Context, companyId uint64, ownerId uint64) error {
	_, err := c.cli.DeleteCompany(ctx, &pb.DeleteCompanyRequest{
		CompanyId: companyId,
		OwnerId:   ownerId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrCompanyNotExists
		case codes.PermissionDenied:
			return model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return model.ErrCoreError
		default:
			return model.ErrCoreUnknown
		}
	}
	return nil
}

func (c *coreClientImpl) GetIndustriesList(ctx context.Context) (map[string]string, error) {
	resp, err := c.cli.GetIndustriesList(ctx, &empty.Empty{})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return map[string]string{}, model.ErrCoreError
		default:
			return map[string]string{}, model.ErrCoreUnknown
		}
	}
	return resp.Data, nil
}

func (c *coreClientImpl) GetIndustryById(ctx context.Context, id uint64) (string, error) {
	resp, err := c.cli.GetIndustryById(ctx, &pb.GetIndustryByIdRequest{Id: id})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return "", model.ErrIndustryNotExists
		case codes.ResourceExhausted:
			return "", model.ErrCoreError
		default:
			return "", model.ErrCoreUnknown
		}
	}
	return resp.Industry, nil
}
