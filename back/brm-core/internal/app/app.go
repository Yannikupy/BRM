package app

import (
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"context"
)

type appImpl struct {
	coreRepo repo.CoreRepo
}

func (a *appImpl) GetCompany(ctx context.Context, id uint) (model.Company, error) {
	return a.coreRepo.GetCompany(ctx, id)
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	return a.coreRepo.CreateCompanyAndOwner(ctx, company, owner)
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd model.UpdateCompany) (model.Company, error) {
	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Company{}, err
	} else if company.OwnerId != ownerId {
		return model.Company{}, model.ErrAuthorization
	}

	return a.coreRepo.UpdateCompany(ctx, companyId, upd)
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) (model.Company, error) {
	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Company{}, err
	} else if company.OwnerId != ownerId {
		return model.Company{}, model.ErrAuthorization
	}

	return a.coreRepo.DeleteCompany(ctx, companyId)
}

func (a *appImpl) CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee model.Employee) (model.Employee, error) {
	if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	} else if company.OwnerId != ownerId {
		return model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.CreateEmployee(ctx, employee)
}

func (a *appImpl) UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd model.UpdateEmployee) (model.Employee, error) {
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	} else if company.OwnerId != ownerId {
		return model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.UpdateEmployee(ctx, employeeId, upd)
}

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (model.Employee, error) {
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	} else if company.OwnerId != ownerId {
		return model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.DeleteEmployee(ctx, employeeId)
}

func (a *appImpl) GetCompanyEmployees(ctx context.Context, companyId uint, employeeId uint, filter model.FilterEmployee) ([]model.Employee, error) {
	_, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.GetCompanyEmployees(ctx, companyId, filter)
}

func (a *appImpl) GetEmployeeByName(ctx context.Context, companyId uint, employeeId uint, ebn model.EmployeeByName) ([]model.Employee, error) {
	_, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.GetEmployeeByName(ctx, ebn)
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint, ownerId uint, employeeId uint) (model.Employee, error) {
	_, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.GetEmployeeById(ctx, employeeId)
}

func (a *appImpl) CreateContact(ctx context.Context, ownerId uint, employeeId uint) (model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	_, err = a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Contact{}, err
	}

	return a.coreRepo.CreateContact(ctx, ownerId, employeeId)
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd model.UpdateContact) (model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.UpdateContact(ctx, contactId, upd)
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint, contactId uint) (model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.DeleteContact(ctx, contactId)
}

func (a *appImpl) GetContacts(ctx context.Context, ownerId uint, pagination model.GetContacts) ([]model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return []model.Contact{}, err
	}

	return a.coreRepo.GetContacts(ctx, ownerId, pagination)
}

func (a *appImpl) GetContactById(ctx context.Context, ownerId uint, contactId uint) (model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.GetContactById(ctx, contactId)
}
