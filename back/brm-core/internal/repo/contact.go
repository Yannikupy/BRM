package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

const (
	createContactQuery = `
		INSERT INTO $1 ("owner_id", "employee_id", "notes", "creation_date", "is_deleted")
		VALUES ($2, $3, $4, $5, $6)
		RETURNING "id";`

	updateContactQuery = `
		UPDATE $1
		SET "notes" = $3
		WHERE "id" = $2 AND (NOT "is_deleted");`

	deleteContactQuery = `
		UPDATE $1
		SET "is_deleted" = true
		WHERE "id" = $2 AND (NOT "is_deleted");`

	// TODO оптимизировать, добавить inner join
	getContactsQuery = `
		SELECT * FROM $1
		WHERE "owner_id" = $2 AND (NOT "is_deleted")
		LIMIT $3 OFFSET $4;`

	getContactByIdQuery = `
		SELECT * FROM $1
		WHERE "id" = $2 AND (NOT "is_deleted");`
)

func (c *coreRepoImpl) CreateContact(ctx context.Context, contact model.Contact) (model.Contact, error) {
	var contactId uint64
	if err := c.QueryRow(ctx, createContactQuery,
		getShardName(contact.OwnerId),
		contact.OwnerId,
		contact.EmployeeId,
		contact.Notes,
		contact.CreationDate,
		contact.IsDeleted,
	).Scan(contactId); err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}

	contact.Id = contactId

	if employee, err := c.GetEmployeeById(ctx, contact.EmployeeId); err != nil {
		return model.Contact{}, err
	} else {
		contact.Empl = employee
	}

	return contact, nil
}

func (c *coreRepoImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error) {
	if e, err := c.Exec(ctx, updateContactQuery,
		getShardName(ownerId),
		contactId,
		upd.Notes,
	); err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Contact{}, model.ErrContactNotExists
	}

	return c.GetContactById(ctx, ownerId, contactId)
}

func (c *coreRepoImpl) DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error {
	if e, err := c.Exec(ctx, deleteContactQuery,
		getShardName(ownerId),
		contactId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrContactNotExists
	} else {
		return nil
	}
}

func (c *coreRepoImpl) GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, error) {
	rows, err := c.Query(ctx, getContactsQuery,
		getShardName(ownerId),
		ownerId,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		return []model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	contacts := make([]model.Contact, 0)
	for rows.Next() {
		var contact model.Contact
		_ = rows.Scan(
			&contact.Id,
			&contact.OwnerId,
			&contact.EmployeeId,
			&contact.Notes,
			&contact.CreationDate,
			&contact.IsDeleted,
		)

		// TODO оптимизировать, добавить inner join
		empl, err := c.GetEmployeeById(ctx, contact.EmployeeId)
		if err != nil {
			return []model.Contact{}, err
		}
		contact.Empl = empl

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (c *coreRepoImpl) GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (model.Contact, error) {
	row := c.QueryRow(ctx, getContactByIdQuery,
		getShardName(ownerId),
		contactId,
	)
	var contact model.Contact
	if err := row.Scan(
		&contact.Id,
		&contact.OwnerId,
		&contact.EmployeeId,
		&contact.Notes,
		&contact.CreationDate,
		&contact.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Contact{}, model.ErrContactNotExists
	} else if err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}

	empl, err := c.GetEmployeeById(ctx, contact.EmployeeId)
	if err != nil {
		return model.Contact{}, err
	}
	contact.Empl = empl

	return contact, nil
}

func getShardName(ownerId uint64) string {
	switch ownerId % 4 {
	case 0:
		return "contact_shard01"
	case 1:
		return "contact_shard02"
	case 2:
		return "contact_shard03"
	case 3:
		return "contact_shard04"
	default:
		// ахуеть как ты вообще попал сюда?
		return ""
	}
}
