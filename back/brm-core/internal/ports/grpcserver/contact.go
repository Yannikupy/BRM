package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *Server) DeleteContact(ctx context.Context, req *pb.DeleteContactRequest) (*empty.Empty, error) {
	// TODO implement
	return &empty.Empty{}, nil
}

func (s *Server) GetContacts(ctx context.Context, req *pb.GetContactsRequest) (*pb.GetContactsResponse, error) {
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

	resp := &pb.GetContactsResponse{
		List: make([]*pb.Contact, len(contacts)),
	}

	for i, contact := range contacts {
		resp.List[i] = &pb.Contact{
			Id:           uint64(contact.Id),
			OwnerId:      uint64(contact.OwnerId),
			EmployeeId:   uint64(contact.EmployeeId),
			Notes:        contact.Notes,
			CreationDate: contact.CreationDate,
			IsDeleted:    contact.IsDeleted,
			Empl: &pb.Employee{
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

func (s *Server) GetContactById(ctx context.Context, req *pb.GetContactByIdRequest) (*pb.GetContactByIdResponse, error) {
	contact, err := s.App.GetContactById(ctx, uint(req.OwnerId), uint(req.ContactId))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetContactByIdResponse{
		Contact: &pb.Contact{
			Id:           uint64(contact.Id),
			OwnerId:      uint64(contact.OwnerId),
			EmployeeId:   uint64(contact.EmployeeId),
			Notes:        contact.Notes,
			CreationDate: contact.CreationDate,
			IsDeleted:    contact.IsDeleted,
			Empl: &pb.Employee{
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
		},
	}, nil
}
