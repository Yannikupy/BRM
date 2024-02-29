package grpccore

import (
	"context"
	"registration/internal/model"
)

type CoreClient interface {
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	GetIndustriesList(ctx context.Context) (map[string]uint64, error)
}
