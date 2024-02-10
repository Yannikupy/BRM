package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) CreateContact(ctx context.Context, ownerId uint, employeeId uint) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}

func (a *appImpl) UpdateContact(ctx context.Context, ownerId uint, contactId uint, upd core.UpdateContact) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}

func (a *appImpl) DeleteContact(ctx context.Context, ownerId uint, contactId uint) error {
	// TODO implement
	return nil
}

func (a *appImpl) GetContacts(ctx context.Context, ownerId uint, pagination core.GetContacts) ([]core.Contact, error) {
	// TODO implement
	return nil, nil
}

func (a *appImpl) GetContactById(ctx context.Context, ownerId uint, contactId uint) (core.Contact, error) {
	// TODO implement
	return core.Contact{}, nil
}
