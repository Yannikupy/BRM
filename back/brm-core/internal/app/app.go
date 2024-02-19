package app

import (
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"context"
	"time"
)

type appImpl struct {
	coreRepo repo.CoreRepo
	auth     grpcauth.AuthClient
}

func (a *appImpl) GetCompany(ctx context.Context, id uint) (model.Company, error) {
	return a.coreRepo.GetCompany(ctx, id)
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	// setting company fields
	company.Id = 0
	company.OwnerId = 0
	company.Rating = 4.
	company.CreationDate = time.Now().UTC()
	company.IsDeleted = false

	// setting owner fields
	owner.Id = 0
	owner.CompanyId = 0
	owner.CreationDate = time.Now().UTC()
	owner.IsDeleted = false

	newCompany, newOwner, err := a.coreRepo.CreateCompanyAndOwner(ctx, company, owner)
	if err != nil {
		return model.Company{}, model.Employee{}, err
	}

	if err = a.auth.RegisterEmployee(ctx, model.EmployeeCredentials{
		Email:      newOwner.Email,
		Password:   newOwner.Password,
		EmployeeId: newOwner.Id,
		CompanyId:  newOwner.CompanyId,
	}); err != nil {
		return model.Company{}, model.Employee{}, err
	}

	return newCompany, newOwner, nil
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint, ownerId uint, upd model.UpdateCompany) (model.Company, error) {
	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Company{}, err
	} else if company.OwnerId != ownerId {
		return model.Company{}, model.ErrAuthorization
	}

	if company.OwnerId != upd.OwnerId {
		var newOwner model.Employee
		newOwner, err = a.coreRepo.GetEmployeeById(ctx, upd.OwnerId)
		if err != nil {
			return model.Company{}, err
		}

		if newOwner.CompanyId != companyId {
			return model.Company{}, model.ErrEmployeeNotExists
		}
	}

	return a.coreRepo.UpdateCompany(ctx, companyId, upd)
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint, ownerId uint) error {
	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return err
	} else if company.OwnerId != ownerId {
		return model.ErrAuthorization
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

	// setting up employee fields
	employee.Id = 0
	employee.CreationDate = time.Now().UTC()
	employee.IsDeleted = false

	newEmployee, err := a.coreRepo.CreateEmployee(ctx, employee)
	if err != nil {
		return model.Employee{}, err
	}

	err = a.auth.RegisterEmployee(ctx, model.EmployeeCredentials{
		Email:      newEmployee.Email,
		Password:   newEmployee.Password,
		EmployeeId: newEmployee.Id,
		CompanyId:  newEmployee.CompanyId,
	})
	if err != nil {
		return model.Employee{}, err
	}

	return newEmployee, nil
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

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error {
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return err
	} else if companyId != employee.CompanyId {
		return model.ErrAuthorization
	}

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return err
	} else if company.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	err = a.coreRepo.DeleteEmployee(ctx, employeeId)
	if err != nil {
		return err
	}
	return a.auth.DeleteEmployee(ctx, employee.Email)
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

	return a.coreRepo.GetEmployeeByName(ctx, companyId, ebn)
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint, _ uint, employeeIdToFind uint) (model.Employee, error) {
	_, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeIdToFind)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	return a.coreRepo.GetEmployeeById(ctx, employeeIdToFind)
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

	return a.coreRepo.CreateContact(ctx, model.Contact{
		Id:           0,
		OwnerId:      ownerId,
		EmployeeId:   employeeId,
		Notes:        "",
		CreationDate: time.Now().UTC(),
		IsDeleted:    false,
		Empl:         model.Employee{},
	})
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd model.UpdateContact) (model.Contact, error) {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.UpdateContact(ctx, ownerId, contactId, upd)
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint, contactId uint) error {
	_, err := a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return err
	} else if contact.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	return a.coreRepo.DeleteContact(ctx, ownerId, contactId)
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

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	return a.coreRepo.GetContactById(ctx, ownerId, contactId)
}
