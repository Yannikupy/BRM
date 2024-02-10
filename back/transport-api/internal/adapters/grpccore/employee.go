package grpccore

import (
	"context"
	"transport-api/internal/model/core"
)

func (c *coreClientImpl) CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee core.Employee) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}

func (c *coreClientImpl) UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd core.UpdateEmployee) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}

func (c *coreClientImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error {
	// TODO implement
	return nil
}

func (c *coreClientImpl) GetCompanyEmployees(ctx context.Context, companyId uint, ownerId uint, filter core.FilterEmployee) ([]core.Employee, error) {
	// TODO implement
	return nil, nil
}

func (c *coreClientImpl) GetEmployeeByName(ctx context.Context, companyId uint, ownerId uint, ebn core.EmployeeByName) ([]core.Employee, error) {
	// TODO implement
	return nil, nil
}

func (c *coreClientImpl) GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}
