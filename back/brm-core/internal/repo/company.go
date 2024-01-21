package repo

import (
	"brm-core/internal/model"
	"context"
)

func (c *coreRepoImpl) GetCompany(ctx context.Context, id uint) (model.Company, error) {
	// TODO: implement
	return model.Company{}, nil
}

func (c *coreRepoImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	// TODO: implement
	return model.Company{}, model.Employee{}, nil
}

func (c *coreRepoImpl) UpdateCompany(ctx context.Context, companyId uint, upd model.UpdateCompany) (model.Company, error) {
	// TODO: implement
	return model.Company{}, nil
}

func (c *coreRepoImpl) DeleteCompany(ctx context.Context, companyId uint) (model.Company, error) {
	// TODO: implement
	return model.Company{}, nil
}
