package app

import (
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"brm-core/pkg/logger"
	"context"
	"errors"
	"time"
)

type appImpl struct {
	coreRepo repo.CoreRepo
	auth     grpcauth.AuthClient

	logs logger.Logger
}

func (a *appImpl) GetCompany(ctx context.Context, id uint64) (model.Company, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": id,
			"Method":     "GetCompany",
		}, err)
	}()
	company, err := a.coreRepo.GetCompany(ctx, id)
	return company, err
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_name": company.Name,
			"owner_email":  owner.Email,
			"Method":       "CreateCompanyAndOwner",
		}, err)
	}()

	if err != nil {
		return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}

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

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd model.UpdateCompany) (model.Company, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"Method":     "UpdateCompany",
		}, err)
	}()

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

	company, err = a.coreRepo.UpdateCompany(ctx, companyId, upd)
	return company, err
}

func (a *appImpl) DeleteCompany(ctx context.Context, companyId uint64, ownerId uint64) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"method":     "DeleteCompany",
		}, err)
	}()

	company, err := a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return err
	} else if company.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	err = a.coreRepo.DeleteCompany(ctx, companyId)
	return err
}

func (a *appImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	return a.coreRepo.GetIndustries(ctx)
}

func (a *appImpl) CreateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employee model.Employee) (model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":         companyId,
			"owner_id":           ownerId,
			"new_employee_email": employee.Email,
			"method":             "CreateEmployee",
		}, err)
	}()

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

func (a *appImpl) UpdateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64, upd model.UpdateEmployee) (model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"method":     "UpdateEmployee",
		}, err)
	}()

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

	employee, err = a.coreRepo.UpdateEmployee(ctx, employeeId, upd)
	return employee, err
}

func (a *appImpl) DeleteEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"owner_id":   ownerId,
			"Method":     "DeleteEmployee",
		}, err)
	}()

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
	err = a.auth.DeleteEmployee(ctx, employee.Email)
	return err
}

func (a *appImpl) GetCompanyEmployees(ctx context.Context, companyId uint64, employeeId uint64, filter model.FilterEmployee) ([]model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetCompanyEmployees",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, model.ErrAuthorization
	}

	employees, err := a.coreRepo.GetCompanyEmployees(ctx, companyId, filter)
	return employees, err
}

func (a *appImpl) GetEmployeeByName(ctx context.Context, companyId uint64, employeeId uint64, ebn model.EmployeeByName) ([]model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetEmployeeByName",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return []model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return []model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return []model.Employee{}, model.ErrAuthorization
	}

	employees, err := a.coreRepo.GetEmployeeByName(ctx, companyId, ebn)
	return employees, err
}

func (a *appImpl) GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (model.Employee, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetEmployeeById",
		}, err)
	}()

	_, err = a.coreRepo.GetCompany(ctx, companyId)
	if err != nil {
		return model.Employee{}, err
	}
	employee, err := a.coreRepo.GetEmployeeById(ctx, employeeIdToFind)
	if err != nil {
		return model.Employee{}, err
	} else if companyId != employee.CompanyId {
		return model.Employee{}, model.ErrAuthorization
	}

	employee, err = a.coreRepo.GetEmployeeById(ctx, employeeIdToFind)
	return employee, err
}

func (a *appImpl) CreateContact(ctx context.Context, ownerId uint64, employeeId uint64) (model.Contact, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id":         ownerId,
			"employee_id_from_contact": employeeId,
			"Method":                   "CreateContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	_, err = a.coreRepo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.CreateContact(ctx, model.Contact{
		Id:           0,
		OwnerId:      ownerId,
		EmployeeId:   employeeId,
		Notes:        "",
		CreationDate: time.Now().UTC(),
		IsDeleted:    false,
		Empl:         model.Employee{},
	})
	return contact, err
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "UpdateContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	contact, err = a.coreRepo.UpdateContact(ctx, ownerId, contactId, upd)
	return contact, err
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "DeleteContact",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return err
	} else if contact.OwnerId != ownerId {
		return model.ErrAuthorization
	}

	err = a.coreRepo.DeleteContact(ctx, ownerId, contactId)
	return err
}

func (a *appImpl) GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"Method":           "GetContacts",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return []model.Contact{}, err
	}

	contacts, err := a.coreRepo.GetContacts(ctx, ownerId, pagination)
	return contacts, err
}

func (a *appImpl) GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (model.Contact, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"contact_owner_id": ownerId,
			"contact_id":       contactId,
			"Method":           "GetContactById",
		}, err)
	}()

	_, err = a.coreRepo.GetEmployeeById(ctx, ownerId)
	if err != nil {
		return model.Contact{}, err
	}

	contact, err := a.coreRepo.GetContactById(ctx, ownerId, contactId)
	if err != nil {
		return model.Contact{}, err
	} else if contact.OwnerId != ownerId {
		return model.Contact{}, err
	}

	contact, err = a.coreRepo.GetContactById(ctx, ownerId, contactId)
	return contact, err
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrAuthServiceError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
