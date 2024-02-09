package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*empty.Empty, error) {
	// TODO implement
	return &empty.Empty{}, nil
}

func (s *Server) GetCompanyEmployees(ctx context.Context, req *pb.GetCompanyEmployeesRequest) (*pb.GetCompanyEmployeesResponse, error) {
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

	resp := &pb.GetCompanyEmployeesResponse{
		List: make([]*pb.Employee, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = &pb.Employee{
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

func (s *Server) GetEmployeeByName(ctx context.Context, req *pb.GetEmployeeByNameRequest) (*pb.GetEmployeeByNameResponse, error) {
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

	resp := &pb.GetEmployeeByNameResponse{
		List: make([]*pb.Employee, len(employees)),
	}
	for i, empl := range employees {
		resp.List[i] = &pb.Employee{
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

func (s *Server) GetEmployeeById(ctx context.Context, req *pb.GetEmployeeByIdRequest) (*pb.GetEmployeeByIdResponse, error) {
	employee, err := s.App.GetEmployeeById(
		ctx,
		uint(req.CompanyId),
		uint(req.OwnerId),
		uint(req.EmployeeId),
	)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetEmployeeByIdResponse{
		Employee: &pb.Employee{
			Id:           uint64(employee.Id),
			CompanyId:    uint64(employee.CompanyId),
			FirstName:    employee.FirstName,
			SecondName:   employee.SecondName,
			Email:        employee.Email,
			JobTitle:     employee.JobTitle,
			Department:   employee.Department,
			CreationDate: employee.CreationDate,
			IsDeleted:    employee.IsDeleted,
		},
	}, nil
}
