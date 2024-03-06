package app

import (
	"context"
	"transport-api/internal/model/ads"
	"transport-api/internal/model/core"
	"transport-api/internal/model/leads"
)

type App interface {
	CoreCompany
	CoreEmployee
	CoreContact
	Ads
	Leads
}

type CoreCompany interface {
	GetCompany(ctx context.Context, id uint64) (core.Company, error)
	UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error)
	DeleteCompany(ctx context.Context, companyId uint64, ownerId uint64) error

	GetIndustriesList(ctx context.Context) (map[string]uint64, error)
	GetIndustryById(ctx context.Context, id uint64) (string, error)
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

type Ads interface {
	GetAdById(ctx context.Context, id uint64) (ads.Ad, error)
	GetAdsList(ctx context.Context, params ads.ListParams) ([]ads.Ad, error)
	CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad ads.Ad) (ads.Ad, error)
	UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd ads.UpdateAd) (ads.Ad, error)
	DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error

	CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (ads.Response, error)
	GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]ads.Response, error)
}

type Leads interface {
	GetLeads(ctx context.Context, companyId uint64, employeeId uint64, filter leads.Filter) ([]leads.Lead, error)
	GetLeadById(ctx context.Context, companyId uint64, employeeId uint64, leadId uint64) (leads.Lead, error)
	UpdateLead(ctx context.Context, companyId uint64, employeeId uint64, id uint64, upd leads.UpdateLead) (leads.Lead, error)

	GetStatuses(ctx context.Context) (map[string]uint64, error)
	GetStatusById(ctx context.Context, id uint64) (string, error)
}
