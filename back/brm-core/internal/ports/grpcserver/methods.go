package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
)

func (s *Server) GetCompany(ctx context.Context, req *pb.GetCompanyRequest) (*pb.CompanyResponse, error) {
	company, err := s.App.GetCompany(ctx, uint(req.Id))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.CompanyResponse{
		Id:           uint64(company.Id),
		Name:         company.Name,
		Description:  company.Description,
		Industry:     uint64(company.Industry),
		OwnerId:      uint64(company.OwnerId),
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}, nil
}

func (s *Server) GetCompanyEmployees(ctx context.Context, req *pb.GetCompanyEmployeesRequest) (*pb.EmployeesListResponse, error) {
	employees, err := s.App.GetCompanyEmployees(
		ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		model.FilterEmployee{
			ByJobTitle:   req.Filter.ByJobTitle,
			JobTitle:     req.Filter.JobTitle,
			ByDepartment: req.Filter.ByDepartment,
			Department:   req.Filter.Department,
			Limit:        int(req.Filter.Limit),
			Offset:       int(req.Filter.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.EmployeesListResponse{
		List: make([]*pb.EmployeeResponse, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = &pb.EmployeeResponse{
			Id:           uint64(empl.Id),
			CompanyId:    uint64(empl.CompanyId),
			FirstName:    empl.FirstName,
			SecondName:   empl.SecondName,
			Email:        empl.Email,
			JobTitle:     empl.JobTitle,
			Department:   empl.Department,
			CreationDate: empl.CreationDate,
			IsDeleted:    empl.IsDeleted,
		}
	}
	return resp, nil
}

func (s *Server) GetEmployeeByName(ctx context.Context, req *pb.GetEmployeeByNameRequest) (*pb.EmployeesListResponse, error) {
	employees, err := s.App.GetEmployeeByName(
		ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		model.EmployeeByName{
			Pattern: req.Ebn.Pattern,
			Limit:   int(req.Ebn.Limit),
			Offset:  int(req.Ebn.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.EmployeesListResponse{
		List: make([]*pb.EmployeeResponse, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = &pb.EmployeeResponse{
			Id:           uint64(empl.Id),
			CompanyId:    uint64(empl.CompanyId),
			FirstName:    empl.FirstName,
			SecondName:   empl.SecondName,
			Email:        empl.Email,
			JobTitle:     empl.JobTitle,
			Department:   empl.Department,
			CreationDate: empl.CreationDate,
			IsDeleted:    empl.IsDeleted,
		}
	}
	return resp, nil
}

func (s *Server) GetEmployeeById(ctx context.Context, req *pb.GetEmployeeByIdRequest) (*pb.EmployeeResponse, error) {
	employee, err := s.App.GetEmployeeById(
		ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		uint(req.EmployeeId),
	)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.EmployeeResponse{
		Id:           uint64(employee.Id),
		CompanyId:    uint64(employee.CompanyId),
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}, nil
}

func (s *Server) GetContacts(ctx context.Context, req *pb.GetContactsRequest) (*pb.ContactsListResponse, error) {
	contacts, err := s.App.GetContacts(
		ctx,
		uint(req.OwnerId),
		model.GetContacts{
			Limit:  int(req.Pagination.Limit),
			Offset: int(req.Pagination.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.ContactsListResponse{
		List: make([]*pb.ContactResponse, len(contacts)),
	}

	for i, contact := range contacts {
		resp.List[i] = &pb.ContactResponse{
			Id:           uint64(contact.Id),
			OwnerId:      uint64(contact.OwnerId),
			EmployeeId:   uint64(contact.EmployeeId),
			Notes:        contact.Notes,
			CreationDate: contact.CreationDate,
			IsDeleted:    contact.IsDeleted,
			Empl: &pb.EmployeeResponse{
				Id:           uint64(contact.Empl.Id),
				CompanyId:    uint64(contact.Empl.CompanyId),
				FirstName:    contact.Empl.FirstName,
				SecondName:   contact.Empl.SecondName,
				Email:        contact.Empl.Email,
				JobTitle:     contact.Empl.JobTitle,
				Department:   contact.Empl.Department,
				CreationDate: contact.Empl.CreationDate,
				IsDeleted:    contact.Empl.IsDeleted,
			},
		}
	}
	return resp, nil
}

func (s *Server) GetContactById(ctx context.Context, req *pb.GetContactByIdRequest) (*pb.ContactResponse, error) {
	contact, err := s.App.GetContactById(ctx, uint(req.OwnerId), uint(req.ContactId))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.ContactResponse{
		Id:           uint64(contact.Id),
		OwnerId:      uint64(contact.OwnerId),
		EmployeeId:   uint64(contact.EmployeeId),
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl: &pb.EmployeeResponse{
			Id:           uint64(contact.Empl.Id),
			CompanyId:    uint64(contact.Empl.CompanyId),
			FirstName:    contact.Empl.FirstName,
			SecondName:   contact.Empl.SecondName,
			Email:        contact.Empl.Email,
			JobTitle:     contact.Empl.JobTitle,
			Department:   contact.Empl.Department,
			CreationDate: contact.Empl.CreationDate,
			IsDeleted:    contact.Empl.IsDeleted,
		},
	}, nil
}
