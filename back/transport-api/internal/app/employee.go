package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee core.Employee) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}

func (a *appImpl) UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd core.UpdateEmployee) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error {
	// TODO implement
	return nil
}

func (a *appImpl) GetCompanyEmployees(ctx context.Context, companyId uint, ownerId uint, filter core.FilterEmployee) ([]core.Employee, error) {
	// TODO implement
	return nil, nil
}

func (a *appImpl) GetEmployeeByName(ctx context.Context, companyId uint, ownerId uint, ebn core.EmployeeByName) ([]core.Employee, error) {
	// TODO implement
	return nil, nil
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (core.Employee, error) {
	// TODO implement
	return core.Employee{}, nil
}
