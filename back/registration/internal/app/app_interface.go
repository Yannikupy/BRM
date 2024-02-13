package app

import (
	"context"
	"registration/internal/model"
)

type App interface {
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
}
