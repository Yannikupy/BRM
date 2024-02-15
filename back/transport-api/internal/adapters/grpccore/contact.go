package grpccore

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToContact(contact *pb.Contact) core.Contact {
	if contact == nil {
		return core.Contact{}
	}
	return core.Contact{
		Id:           uint(contact.Id),
		OwnerId:      uint(contact.OwnerId),
		EmployeeId:   uint(contact.EmployeeId),
		Notes:        contact.Notes,
		CreationDate: contact.CreationDate,
		IsDeleted:    contact.IsDeleted,
		Empl:         respToEmployee(contact.Empl),
	}
}

func (c *coreClientImpl) CreateContact(ctx context.Context, ownerId uint, employeeId uint) (core.Contact, error) {
	resp, err := c.cli.CreateContact(ctx, &pb.CreateContactRequest{
		OwnerId:    uint64(ownerId),
		EmployeeId: uint64(employeeId),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Contact{}, model.ErrEmployeeNotExists
		case codes.ResourceExhausted:
			return core.Contact{}, model.ErrCoreError
		default:
			return core.Contact{}, model.ErrCoreUnknown
		}
	}
	return respToContact(resp.Contact), nil
}

func (c *coreClientImpl) UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd core.UpdateContact) (core.Contact, error) {
	resp, err := c.cli.UpdateContact(ctx, &pb.UpdateContactRequest{
		OwnerId:   uint64(ownerId),
		ContactId: uint64(contactId),
		Upd:       &pb.UpdateContactFields{Notes: upd.Notes},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Contact{}, model.ErrContactNotExists
		case codes.PermissionDenied:
			return core.Contact{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Contact{}, model.ErrCoreError
		default:
			return core.Contact{}, model.ErrCoreUnknown
		}
	}
	return respToContact(resp.Contact), nil
}

func (c *coreClientImpl) DeleteContact(ctx context.Context, ownerId uint, contactId uint) error {
	_, err := c.cli.DeleteContact(ctx, &pb.DeleteContactRequest{
		OwnerId:   uint64(ownerId),
		ContactId: uint64(contactId),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrContactNotExists
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

func (c *coreClientImpl) GetContacts(ctx context.Context, ownerId uint, pagination core.GetContacts) ([]core.Contact, error) {
	resp, err := c.cli.GetContacts(ctx, &pb.GetContactsRequest{
		OwnerId: uint64(ownerId),
		Pagination: &pb.GetContactsPagination{
			Limit:  int64(pagination.Limit),
			Offset: int64(pagination.Offset),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return []core.Contact{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return []core.Contact{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []core.Contact{}, model.ErrCoreError
		default:
			return []core.Contact{}, model.ErrCoreUnknown
		}
	}
	contacts := make([]core.Contact, len(resp.List))
	for i, contact := range resp.List {
		contacts[i] = respToContact(contact)
	}
	return contacts, nil
}

func (c *coreClientImpl) GetContactById(ctx context.Context, ownerId uint, contactId uint) (core.Contact, error) {
	resp, err := c.cli.GetContactById(ctx, &pb.GetContactByIdRequest{
		OwnerId:   uint64(ownerId),
		ContactId: uint64(contactId),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			// костыль, ну а чё ещё поделать
			if strings.Contains(err.Error(), "contact") {
				return core.Contact{}, model.ErrContactNotExists
			} else {
				return core.Contact{}, model.ErrEmployeeNotExists
			}
		case codes.PermissionDenied:
			return core.Contact{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Contact{}, model.ErrCoreError
		default:
			return core.Contact{}, model.ErrCoreUnknown
		}
	}
	return respToContact(resp.Contact), nil
}
