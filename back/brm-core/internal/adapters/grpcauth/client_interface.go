package grpcauth

import (
	"brm-core/internal/model"
	"context"
)

type AuthClient interface {
	RegisterEmployee(ctx context.Context, creds model.EmployeeCredentials) error
	DeleteEmployee(ctx context.Context, email string) error
}
