package repo

import (
	"brm-core/internal/model"
	"context"
)

type CompanyRepo interface {
	GetCompany(ctx context.Context, id uint) (model.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint, upd model.UpdateCompany) (model.Company, error)
	DeleteCompany(ctx context.Context, companyId uint) (model.Company, error)
}

type EmployeeRepo interface {
	CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, employeeId uint, upd model.UpdateEmployee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, employeeId uint) (model.Employee, error)

	GetCompanyEmployees(ctx context.Context, companyId uint, filter model.FilterEmployee) ([]model.Employee, error)
	GetEmployeeByName(ctx context.Context, ebn model.EmployeeByName) ([]model.Employee, error)
	GetEmployeeById(ctx context.Context, employeeId uint) (model.Employee, error)
}

type ContactRepo interface {
	CreateContact(ctx context.Context, ownerId uint, employeeId uint) (model.Contact, error)
	UpdateContact(ctx context.Context, contactId uint, upd model.UpdateContact) (model.Contact, error)
	DeleteContact(ctx context.Context, contactId uint) (model.Contact, error)

	GetContacts(ctx context.Context, ownerId uint, pagination model.GetContacts) ([]model.Contact, error)
	GetContactById(ctx context.Context, contactId uint) (model.Contact, error)
}
