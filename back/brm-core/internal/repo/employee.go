package repo

import (
	"brm-core/internal/model"
	"context"
)

func (c *coreRepoImpl) CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error) {
	// TODO: implement
	return model.Employee{}, nil
}

func (c *coreRepoImpl) UpdateEmployee(ctx context.Context, employeeId uint, upd model.UpdateEmployee) (model.Employee, error) {
	// TODO: implement
	return model.Employee{}, nil
}

func (c *coreRepoImpl) DeleteEmployee(ctx context.Context, employeeId uint) (model.Employee, error) {
	// TODO: implement
	return model.Employee{}, nil
}

func (c *coreRepoImpl) GetCompanyEmployees(ctx context.Context, companyId uint, filter model.FilterEmployee) ([]model.Employee, error) {
	// TODO: implement
	return []model.Employee{}, nil
}

func (c *coreRepoImpl) GetEmployeeByName(ctx context.Context, ebn model.EmployeeByName) ([]model.Employee, error) {
	// TODO: implement
	return []model.Employee{}, nil
}

func (c *coreRepoImpl) GetEmployeeById(ctx context.Context, employeeId uint) (model.Employee, error) {
	// TODO: implement
	return model.Employee{}, nil
}
