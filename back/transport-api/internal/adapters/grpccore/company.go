package grpccore

import (
	"context"
	"transport-api/internal/model/core"
)

func (c *coreClientImpl) GetCompany(ctx context.Context, id uint) (core.Company, error) {
	// TODO implement
	return core.Company{}, nil
}

func (c *coreClientImpl) CreateCompanyAndOwner(ctx context.Context, company core.Company, owner core.Employee) (core.Company, core.Employee, error) {
	// TODO implement
	return core.Company{}, core.Employee{}, nil
}

func (c *coreClientImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd core.UpdateCompany) (core.Company, error) {
	// TODO implement
	return core.Company{}, nil
}

func (c *coreClientImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error {
	// TODO implement
	return nil
}
