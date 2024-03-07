package grpccore

import (
	"context"
	"transport-api/internal/model/core"
)

type CoreClient interface {
	CoreCompany
	CoreEmployee
	CoreContact
}

type CoreCompany interface {
	GetCompany(ctx context.Context, id uint64) (core.Company, error)
	UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error)
	DeleteCompany(ctx context.Context, companyId uint64, ownerId uint64) error

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}

type CoreEmployee interface {
	CreateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employee core.Employee) (core.Employee, error)
	UpdateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64, upd core.UpdateEmployee) (core.Employee, error)
	DeleteEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64) error
	GetCompanyEmployees(ctx context.Context, companyId uint64, employeeId uint64, filter core.FilterEmployee) ([]core.Employee, error)
	GetEmployeeByName(ctx context.Context, companyId uint64, employeeId uint64, ebn core.EmployeeByName) ([]core.Employee, error)
	GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (core.Employee, error)
}

type CoreContact interface {
	CreateContact(ctx context.Context, ownerId uint64, employeeId uint64) (core.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd core.UpdateContact) (core.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error
	GetContacts(ctx context.Context, ownerId uint64, pagination core.GetContacts) ([]core.Contact, error)
	GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (core.Contact, error)
}
