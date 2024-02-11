package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func employeeToModelEmployee(employee *pb.Employee) model.Employee {
	if employee == nil {
		return model.Employee{}
	}
	return model.Employee{
		Id:           uint(employee.Id),
		CompanyId:    uint(employee.CompanyId),
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func modelEmployeeToEmployee(employee model.Employee) *pb.Employee {
	if employee.Id == 0 {
		return nil
	}
	return &pb.Employee{
		Id:           uint64(employee.Id),
		CompanyId:    uint64(employee.CompanyId),
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func (s *Server) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	employee, err := s.App.CreateEmployee(ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		employeeToModelEmployee(req.Employee),
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateEmployeeResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}

func (s *Server) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	employee, err := s.App.UpdateEmployee(ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		uint(req.EmployeeId),
		model.UpdateEmployee{
			FirstName:  req.Upd.FirstName,
			SecondName: req.Upd.SecondName,
			Email:      req.Upd.Email,
			JobTitle:   req.Upd.JobTitle,
			Department: req.Upd.Department,
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateEmployeeResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}

func (s *Server) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*empty.Empty, error) {
	if err := s.App.DeleteEmployee(ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		uint(req.EmployeeId),
	); err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetCompanyEmployees(ctx context.Context, req *pb.GetCompanyEmployeesRequest) (*pb.GetCompanyEmployeesResponse, error) {
	employees, err := s.App.GetCompanyEmployees(
		ctx,
		uint(req.CompanyId),
		uint(req.EmployeeId),
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

	resp := &pb.GetCompanyEmployeesResponse{
		List: make([]*pb.Employee, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = modelEmployeeToEmployee(empl)
	}
	return resp, nil
}

func (s *Server) GetEmployeeByName(ctx context.Context, req *pb.GetEmployeeByNameRequest) (*pb.GetEmployeeByNameResponse, error) {
	employees, err := s.App.GetEmployeeByName(
		ctx,
		uint(req.CompanyId),
		uint(req.EmployeeId),
		model.EmployeeByName{
			Pattern: req.Ebn.Pattern,
			Limit:   int(req.Ebn.Limit),
			Offset:  int(req.Ebn.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.GetEmployeeByNameResponse{
		List: make([]*pb.Employee, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = modelEmployeeToEmployee(empl)
	}
	return resp, nil
}

func (s *Server) GetEmployeeById(ctx context.Context, req *pb.GetEmployeeByIdRequest) (*pb.GetEmployeeByIdResponse, error) {
	employee, err := s.App.GetEmployeeById(
		ctx,
		uint(req.CompanyId),
		uint(req.EmployeeId),
		uint(req.EmployeeIdToFind),
	)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetEmployeeByIdResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}
