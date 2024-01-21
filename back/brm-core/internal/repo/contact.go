package repo

import (
	"brm-core/internal/model"
	"context"
)

func (c *coreRepoImpl) CreateContact(ctx context.Context, ownerId uint, employeeId uint) (model.Contact, error) {
	// TODO implement
	return model.Contact{}, nil
}

func (c *coreRepoImpl) UpdateContact(ctx context.Context, contactId uint, upd model.UpdateContact) (model.Contact, error) {
	// TODO implement
	return model.Contact{}, nil
}

func (c *coreRepoImpl) DeleteContact(ctx context.Context, contactId uint) (model.Contact, error) {
	// TODO implement
	return model.Contact{}, nil
}

func (c *coreRepoImpl) GetContacts(ctx context.Context, ownerId uint, pagination model.GetContacts) ([]model.Contact, error) {
	// TODO implement
	return []model.Contact{}, nil
}

func (c *coreRepoImpl) GetContactById(ctx context.Context, contactId uint) (model.Contact, error) {
	// TODO implement
	return model.Contact{}, nil
}
