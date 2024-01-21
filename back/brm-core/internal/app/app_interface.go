package app

import (
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"context"
)

type App interface {
	CompanyApp
	EmployeeApp
	ContactApp
}

func New(
	companyRepo repo.CompanyRepo,
	contactRepo repo.ContactRepo,
	employeeRepo repo.EmployeeRepo,
) App {
	return &appImpl{
		companyRepo:  companyRepo,
		contactRepo:  contactRepo,
		employeeRepo: employeeRepo,
	}
}

type CompanyApp interface {
	GetCompany(ctx context.Context, id uint) (model.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd model.UpdateCompany) (model.Company, error)
	DeleteCompany(ctx context.Context, companyId uint, ownerId uint) (model.Company, error)
}

type EmployeeApp interface {
	CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd model.UpdateEmployee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (model.Employee, error)

	GetCompanyEmployees(ctx context.Context, companyId uint, ownerId uint, filter model.FilterEmployee) ([]model.Employee, error)
	GetEmployeeByName(ctx context.Context, companyId uint, ownerId uint, ebn model.EmployeeByName) ([]model.Employee, error)
	GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (model.Employee, error)
}

type ContactApp interface {
	CreateContact(ctx context.Context, ownerId uint, employeeId uint) (model.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd model.UpdateContact) (model.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint, contactId uint) (model.Contact, error)

	GetContacts(ctx context.Context, ownerId uint, pagination model.GetContacts) ([]model.Contact, error)
	GetContactById(ctx context.Context, ownerId uint, contactId uint) (model.Contact, error)
}
