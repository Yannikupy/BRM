package app

import (
	"context"
	"transport-api/internal/model/core"
)

type App interface {
	CoreCompany
	CoreEmployee
	CoreContact
}

type CoreCompany interface {
	GetCompany(ctx context.Context, id uint) (core.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company core.Company, owner core.Employee) (core.Company, core.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd core.UpdateCompany) (core.Company, error)
	DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error
}

type CoreEmployee interface {
	CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee core.Employee) (core.Employee, error)
	UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd core.UpdateEmployee) (core.Employee, error)
	DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error
	GetCompanyEmployees(ctx context.Context, companyId uint, ownerId uint, filter core.FilterEmployee) ([]core.Employee, error)
	GetEmployeeByName(ctx context.Context, companyId uint, ownerId uint, ebn core.EmployeeByName) ([]core.Employee, error)
	GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (core.Employee, error)
}

type CoreContact interface {
	CreateContact(ctx context.Context, ownerId uint, employeeId uint) (core.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd core.UpdateContact) (core.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint, contactId uint) error
	GetContacts(ctx context.Context, ownerId uint, pagination core.GetContacts) ([]core.Contact, error)
	GetContactById(ctx context.Context, ownerId uint, contactId uint) (core.Contact, error)
}
