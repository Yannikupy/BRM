package repo

import (
	"brm-core/internal/model"
	"context"

	"github.com/jackc/pgx/v5"
)

type CoreRepo interface {
	CompanyRepo
	EmployeeRepo
	ContactRepo
}

func New(conn *pgx.Conn) CoreRepo {
	return &coreRepoImpl{
		Conn: *conn,
	}
}

type CompanyRepo interface {
	GetCompany(ctx context.Context, id uint) (model.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint, upd model.UpdateCompany) (model.Company, error)
	DeleteCompany(ctx context.Context, companyId uint) error
}

type EmployeeRepo interface {
	CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, employeeId uint, upd model.UpdateEmployee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, employeeId uint) error

	GetCompanyEmployees(ctx context.Context, companyId uint, filter model.FilterEmployee) ([]model.Employee, error)
	GetEmployeeByName(ctx context.Context, companyId uint, ebn model.EmployeeByName) ([]model.Employee, error)
	GetEmployeeById(ctx context.Context, employeeId uint) (model.Employee, error)
}

type ContactRepo interface {
	CreateContact(ctx context.Context, contact model.Contact) (model.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd model.UpdateContact) (model.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint, contactId uint) error

	GetContacts(ctx context.Context, ownerId uint, pagination model.GetContacts) ([]model.Contact, error)
	GetContactById(ctx context.Context, ownerId uint, contactId uint) (model.Contact, error)
}
