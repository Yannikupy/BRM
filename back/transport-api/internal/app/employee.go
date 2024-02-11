package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee core.Employee) (core.Employee, error) {
	return a.core.CreateEmployee(ctx, companyId, ownerId, employee)
}

func (a *appImpl) UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd core.UpdateEmployee) (core.Employee, error) {
	return a.core.UpdateEmployee(ctx, companyId, ownerId, employeeId, upd)
}

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error {
	return a.core.DeleteEmployee(ctx, companyId, ownerId, employeeId)
}

func (a *appImpl) GetCompanyEmployees(ctx context.Context, companyId uint, employeeId uint, filter core.FilterEmployee) ([]core.Employee, error) {
	return a.core.GetCompanyEmployees(ctx, companyId, employeeId, filter)
}

func (a *appImpl) GetEmployeeByName(ctx context.Context, companyId uint, ownerId uint, ebn core.EmployeeByName) ([]core.Employee, error) {
	return a.core.GetEmployeeByName(ctx, companyId, ownerId, ebn)
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (core.Employee, error) {
	return a.core.GetEmployeeById(ctx, companyId, ownerId, employeeId)
}
