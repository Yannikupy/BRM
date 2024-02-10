package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) GetCompany(ctx context.Context, id uint) (core.Company, error) {
	// TODO implement
	return core.Company{}, nil
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company core.Company, owner core.Employee) (core.Company, core.Employee, error) {
	// TODO implement
	return core.Company{}, core.Employee{}, nil
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd core.UpdateCompany) (core.Company, error) {
	// TODO implement
	return core.Company{}, nil
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error {
	// TODO implement
	return nil
}
