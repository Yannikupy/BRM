package rmq

import (
	"brm-core/internal/app"
	"brm-core/internal/model"
	"context"
	"encoding/json"
	"fmt"
)

// HandleJobs слушает очередь rabbitmq, запускать в отдельной горутине
func (s *Shard) HandleJobs(ctx context.Context, a app.App) {
	for jobDelivery := range s.consumer.messages {
		var job jobRequest
		err := json.Unmarshal(jobDelivery.Body, &job)
		if err != nil {
			_ = s.producer.publish(jobDelivery.MessageId, jobResponse{
				JobType:   "",
				JobResult: []byte{},
				Error:     fmt.Errorf("unmarshalling job: %w", err).Error(),
			})
			continue
		}

		switch job.JobType {
		case createCompanyAndOwner:
			s.handleCreateCompanyAndOwnerJob(ctx, a, jobDelivery.MessageId, job)
		case updateCompany:
			s.handleUpdateCompanyJob(ctx, a, jobDelivery.MessageId, job)
		case deleteCompany:
			s.handleDeleteCompanyJob(ctx, a, jobDelivery.MessageId, job)
		case createEmployee:
			s.handleCreateEmployeeJob(ctx, a, jobDelivery.MessageId, job)
		case updateEmployee:
			s.handleUpdateEmployeeJob(ctx, a, jobDelivery.MessageId, job)
		case deleteEmployee:
			s.handleDeleteEmployeeJob(ctx, a, jobDelivery.MessageId, job)
		case createContact:
			s.handleCreateContactJob(ctx, a, jobDelivery.MessageId, job)
		case updateContact:
			s.handleUpdateContactJob(ctx, a, jobDelivery.MessageId, job)
		case deleteContact:
			s.handleDeleteContactJob(ctx, a, jobDelivery.MessageId, job)
		default: // unknown
			_ = s.producer.publish(jobDelivery.MessageId, jobResponse{
				JobType:   "",
				JobResult: []byte{},
				Error:     "unknown job type",
			})
		}
	}
}

func (s *Shard) handleCreateCompanyAndOwnerJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData createCompanyAndOwnerJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	company, employee, err := a.CreateCompanyAndOwner(
		ctx,
		model.Company{
			Id:           jobData.Company.Id,
			Name:         jobData.Company.Name,
			Description:  jobData.Company.Description,
			Industry:     jobData.Company.Industry,
			OwnerId:      jobData.Company.OwnerId,
			Rating:       jobData.Company.Rating,
			CreationDate: jobData.Company.CreationDate,
			IsDeleted:    jobData.Company.IsDeleted,
		},
		model.Employee{
			Id:           jobData.Owner.Id,
			CompanyId:    jobData.Owner.CompanyId,
			FirstName:    jobData.Owner.FirstName,
			SecondName:   jobData.Owner.SecondName,
			Email:        jobData.Owner.Email,
			JobTitle:     jobData.Owner.JobTitle,
			Department:   jobData.Owner.Department,
			CreationDate: jobData.Owner.CreationDate,
			IsDeleted:    jobData.Owner.IsDeleted,
		})
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := createCompanyAndOwnerResponse{
		Company: companyData{
			Id:           company.Id,
			Name:         company.Name,
			Description:  company.Description,
			Industry:     company.Industry,
			OwnerId:      company.OwnerId,
			Rating:       company.Rating,
			CreationDate: company.CreationDate,
			IsDeleted:    company.IsDeleted,
		},
		Owner: employeeData{
			Id:           employee.Id,
			CompanyId:    employee.CompanyId,
			FirstName:    employee.FirstName,
			SecondName:   employee.SecondName,
			Email:        employee.Email,
			JobTitle:     employee.JobTitle,
			Department:   employee.Department,
			CreationDate: employee.CreationDate,
			IsDeleted:    employee.IsDeleted,
		},
	}

	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleUpdateCompanyJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData updateCompanyJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	company, err := a.UpdateCompany(
		ctx,
		jobData.CompanyId,
		jobData.OwnerId,
		model.UpdateCompany{
			Name:        jobData.Upd.Name,
			Description: jobData.Upd.Description,
			Industry:    jobData.Upd.Industry,
			OwnerId:     jobData.Upd.OwnerId,
		})
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := updateCompanyResponse{Company: companyData{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}}

	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleDeleteCompanyJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData deleteCompanyJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	err = a.DeleteCompany(
		ctx,
		jobData.CompanyId,
		jobData.OwnerId,
	)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: []byte{},
		Error:     "",
	})
}

func (s *Shard) handleCreateEmployeeJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData createEmployeeJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	employee, err := a.CreateEmployee(
		ctx,
		jobData.CompanyId,
		jobData.OwnerId,
		model.Employee{
			Id:           jobData.Employee.Id,
			CompanyId:    jobData.Employee.CompanyId,
			FirstName:    jobData.Employee.FirstName,
			SecondName:   jobData.Employee.SecondName,
			Email:        jobData.Employee.Email,
			JobTitle:     jobData.Employee.JobTitle,
			Department:   jobData.Employee.Department,
			CreationDate: jobData.Employee.CreationDate,
			IsDeleted:    jobData.Employee.IsDeleted,
		})
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := createEmployeeResponse{Employee: employeeData{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}}

	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleUpdateEmployeeJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData updateEmployeeJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	employee, err := a.UpdateEmployee(
		ctx,
		jobData.CompanyId,
		jobData.OwnerId,
		jobData.EmployeeId,
		model.UpdateEmployee{
			FirstName:  jobData.Upd.FirstName,
			SecondName: jobData.Upd.SecondName,
			Email:      jobData.Upd.Email,
			JobTitle:   jobData.Upd.JobTitle,
			Department: jobData.Upd.Department,
		})
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := updateEmployeeResponse{Employee: employeeData{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}}

	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleDeleteEmployeeJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData deleteEmployeeJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	err = a.DeleteEmployee(
		ctx,
		jobData.CompanyId,
		jobData.OwnerId,
		jobData.EmployeeId,
	)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: []byte{},
		Error:     "",
	})
}

func (s *Shard) handleCreateContactJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData createContactJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	contact, err := a.CreateContact(
		ctx,
		jobData.OwnerId,
		jobData.EmployeeId,
	)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := createContactResponse{Contact: contactData{
		Id:           contact.Id,
		OwnerId:      contact.OwnerId,
		EmployeeId:   contact.EmployeeId,
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl: employeeData{
			Id:           contact.Empl.Id,
			CompanyId:    contact.Empl.CompanyId,
			FirstName:    contact.Empl.FirstName,
			SecondName:   contact.Empl.SecondName,
			Email:        contact.Empl.Email,
			JobTitle:     contact.Empl.JobTitle,
			Department:   contact.Empl.Department,
			CreationDate: contact.Empl.CreationDate,
			IsDeleted:    contact.Empl.IsDeleted,
		},
	}}
	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleUpdateContactJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData updateContactJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	contact, err := a.UpdateContact(
		ctx,
		jobData.OwnerId,
		jobData.ContactId,
		model.UpdateContact{
			Notes: jobData.Upd.Notes,
		})
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	result := updateContactResponse{Contact: contactData{
		Id:           contact.Id,
		OwnerId:      contact.OwnerId,
		EmployeeId:   contact.EmployeeId,
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl: employeeData{
			Id:           contact.Empl.Id,
			CompanyId:    contact.Empl.CompanyId,
			FirstName:    contact.Empl.FirstName,
			SecondName:   contact.Empl.SecondName,
			Email:        contact.Empl.Email,
			JobTitle:     contact.Empl.JobTitle,
			Department:   contact.Empl.Department,
			CreationDate: contact.Empl.CreationDate,
			IsDeleted:    contact.Empl.IsDeleted,
		},
	}}
	resultData, err := json.Marshal(result)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("result marshalling to json: %w", err).Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: resultData,
		Error:     "",
	})
}

func (s *Shard) handleDeleteContactJob(ctx context.Context, a app.App, id string, job jobRequest) {
	var jobData deleteContactJob
	err := json.Unmarshal(job.JobBody, &jobData)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     fmt.Errorf("unmarshalling job body: %w", err).Error(),
		})
		return
	}

	err = a.DeleteContact(
		ctx,
		jobData.OwnerId,
		jobData.ContactId,
	)
	if err != nil {
		_ = s.producer.publish(id, jobResponse{
			JobType:   job.JobType,
			JobResult: []byte{},
			Error:     err.Error(),
		})
		return
	}

	_ = s.producer.publish(id, jobResponse{
		JobType:   job.JobType,
		JobResult: []byte{},
		Error:     "",
	})
}
