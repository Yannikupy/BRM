package grpccore

import "context"

type CoreClient interface {
	GetCompany(ctx context.Context, id uint64) (uint64, error)
	GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (uint64, uint64, error)
}
