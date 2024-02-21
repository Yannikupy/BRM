package grpccore

import "context"

type CoreClient interface {
	// CheckEmployee returns true, nil if employee with id employeeId
	// belongs to company with id companyId, or false and error if not
	CheckEmployee(ctx context.Context, companyId uint64, employeeId uint64) (bool, error)
}
