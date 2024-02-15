package grpccore

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToEmployee(employee *pb.Employee) core.Employee {
	if employee == nil {
		return core.Employee{}
	}
	return core.Employee{
		Id:           uint(employee.Id),
		CompanyId:    uint(employee.CompanyId),
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		Passport:     employee.Password,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func employeeToRequest(employee core.Employee) *pb.Employee {
	return &pb.Employee{
		Id:           uint64(employee.Id),
		CompanyId:    uint64(employee.CompanyId),
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		Password:     employee.Passport,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func (c *coreClientImpl) CreateEmployee(ctx context.Context, companyId uint, ownerId uint, employee core.Employee) (core.Employee, error) {
	resp, err := c.cli.CreateEmployee(ctx, &pb.CreateEmployeeRequest{
		CompanyId: uint64(companyId),
		OwnerId:   uint64(ownerId),
		Employee:  employeeToRequest(employee),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	return respToEmployee(resp.Employee), nil
}

func (c *coreClientImpl) UpdateEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint, upd core.UpdateEmployee) (core.Employee, error) {
	resp, err := c.cli.UpdateEmployee(ctx, &pb.UpdateEmployeeRequest{
		CompanyId:  uint64(companyId),
		OwnerId:    uint64(ownerId),
		EmployeeId: uint64(employeeId),
		Upd: &pb.UpdateEmployeeFields{
			FirstName:  upd.FirstName,
			SecondName: upd.SecondName,
			JobTitle:   upd.JobTitle,
			Department: upd.Department,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	return respToEmployee(resp.Employee), nil
}

func (c *coreClientImpl) DeleteEmployee(ctx context.Context, companyId uint, ownerId uint, employeeId uint) error {
	_, err := c.cli.DeleteEmployee(ctx, &pb.DeleteEmployeeRequest{
		CompanyId:  uint64(companyId),
		OwnerId:    uint64(ownerId),
		EmployeeId: uint64(employeeId),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return model.ErrCoreError
		default:
			return model.ErrCoreUnknown
		}
	}
	return nil
}

func (c *coreClientImpl) GetCompanyEmployees(ctx context.Context, companyId uint, employeeId uint, filter core.FilterEmployee) ([]core.Employee, error) {
	resp, err := c.cli.GetCompanyEmployees(ctx, &pb.GetCompanyEmployeesRequest{
		CompanyId:  uint64(companyId),
		EmployeeId: uint64(employeeId),
		Filter: &pb.FilterEmployee{
			ByJobTitle:   filter.ByJobTitle,
			JobTitle:     filter.JobTitle,
			ByDepartment: filter.ByDepartment,
			Department:   filter.Department,
			Limit:        int64(filter.Limit),
			Offset:       int64(filter.Offset),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return []core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return []core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []core.Employee{}, model.ErrCoreError
		default:
			return []core.Employee{}, model.ErrCoreUnknown
		}
	}
	employees := make([]core.Employee, len(resp.List))
	for i, empl := range resp.List {
		employees[i] = respToEmployee(empl)
	}
	return employees, nil
}

func (c *coreClientImpl) GetEmployeeByName(ctx context.Context, companyId uint, employeeId uint, ebn core.EmployeeByName) ([]core.Employee, error) {
	resp, err := c.cli.GetEmployeeByName(ctx, &pb.GetEmployeeByNameRequest{
		CompanyId:  uint64(companyId),
		EmployeeId: uint64(employeeId),
		Ebn: &pb.EmployeeByName{
			Pattern: ebn.Pattern,
			Limit:   int64(ebn.Limit),
			Offset:  int64(ebn.Offset),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return []core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return []core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []core.Employee{}, model.ErrCoreError
		default:
			return []core.Employee{}, model.ErrCoreUnknown
		}
	}
	employees := make([]core.Employee, len(resp.List))
	for i, empl := range resp.List {
		employees[i] = respToEmployee(empl)
	}
	return employees, nil
}

func (c *coreClientImpl) GetEmployeeById(ctx context.Context, companyId uint, employeeId uint, employeeIdToFind uint) (core.Employee, error) {
	resp, err := c.cli.GetEmployeeById(ctx, &pb.GetEmployeeByIdRequest{
		CompanyId:        uint64(companyId),
		EmployeeId:       uint64(employeeId),
		EmployeeIdToFind: uint64(employeeIdToFind),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	return respToEmployee(resp.Employee), nil
}
