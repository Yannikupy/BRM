package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func contactToModelContact(contact *pb.Contact) model.Contact {
	if contact == nil {
		return model.Contact{}
	}
	return model.Contact{
		Id:           uint(contact.Id),
		OwnerId:      uint(contact.OwnerId),
		EmployeeId:   uint(contact.EmployeeId),
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl:         employeeToModelEmployee(contact.Empl),
	}
}

func modelContactToContact(contact model.Contact) *pb.Contact {
	if contact.Id == 0 {
		return nil
	}
	return &pb.Contact{
		Id:           uint64(contact.Id),
		OwnerId:      uint64(contact.OwnerId),
		EmployeeId:   uint64(contact.EmployeeId),
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl:         modelEmployeeToEmployee(contact.Empl),
	}
}

func (s *Server) CreateContact(ctx context.Context, req *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	contact, err := s.App.CreateContact(ctx,
		uint(req.OwnerId),
		uint(req.EmployeeId),
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateContactResponse{
		Contact: modelContactToContact(contact),
	}, nil
}

func (s *Server) UpdateContact(ctx context.Context, req *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	contact, err := s.App.UpdateContact(ctx,
		uint(req.OwnerId),
		uint(req.ContactId),
		model.UpdateContact{
			Notes: req.Upd.Notes,
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateContactResponse{
		Contact: modelContactToContact(contact),
	}, nil
}

func (s *Server) DeleteContact(ctx context.Context, req *pb.DeleteContactRequest) (*empty.Empty, error) {
	if err := s.App.DeleteContact(ctx,
		uint(req.OwnerId),
		uint(req.ContactId),
	); err != nil {
		return nil, mapErrors(err)
	}
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
		resp.List[i] = modelContactToContact(contact)
	}
	return resp, nil
}

func (s *Server) GetContactById(ctx context.Context, req *pb.GetContactByIdRequest) (*pb.GetContactByIdResponse, error) {
	contact, err := s.App.GetContactById(ctx, uint(req.OwnerId), uint(req.ContactId))
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetContactByIdResponse{
		Contact: modelContactToContact(contact),
	}, nil
}
