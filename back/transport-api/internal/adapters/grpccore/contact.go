package grpccore

import (
	"context"
	"transport-api/internal/model/core"
)

func (c *coreClientImpl) CreateContact(ctx context.Context, ownerId uint, employeeId uint) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}

func (c *coreClientImpl) UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd core.UpdateContact) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}

func (c *coreClientImpl) DeleteContact(ctx context.Context, ownerId uint, contactId uint) error {
	// TODO implement
	return nil
}

func (c *coreClientImpl) GetContacts(ctx context.Context, ownerId uint, pagination core.GetContacts) ([]core.Contact, error) {
	// TODO implement
	return nil, nil
}

func (c *coreClientImpl) GetContactById(ctx context.Context, ownerId uint, contactId uint) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}
