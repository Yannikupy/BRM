package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) GetCompany(ctx context.Context, id uint64) (core.Company, error) {
	return a.core.GetCompany(ctx, id)
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error) {
	return a.core.UpdateCompany(ctx, companyId, ownerId, upd)
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint64, ownerId uint64) error {
	return a.core.DeleteCompany(ctx, companyId, ownerId)
}

func (a *appImpl) GetIndustriesList(ctx context.Context) (map[string]uint64, error) {
	return a.core.GetIndustriesList(ctx)
}

func (a *appImpl) GetIndustryById(ctx context.Context, id uint64) (string, error) {
	return a.core.GetIndustryById(ctx, id)
}
