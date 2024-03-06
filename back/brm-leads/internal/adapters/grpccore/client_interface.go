package grpccore

import "context"

type CoreClient interface {
	GetCompanyName(ctx context.Context, id uint64) (string, error)
	GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (uint64, uint64, error)
}
