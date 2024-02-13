package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) GetCompany(ctx context.Context, id uint) (core.Company, error) {
	return a.core.GetCompany(ctx, id)
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd core.UpdateCompany) (core.Company, error) {
	return a.core.UpdateCompany(ctx, companyId, ownerId, upd)
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error {
	return a.core.DeleteCompany(ctx, companyId, ownerId)
}
