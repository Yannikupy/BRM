package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (c *coreRepoImpl) CreateContact(ctx context.Context, contact model.Contact) (model.Contact, error) {
	var contactId uint64
	var pgErr *pgconn.PgError
	if err := c.QueryRow(ctx, getCreateContactQuery(contact.OwnerId),
		contact.OwnerId,
		contact.EmployeeId,
		contact.Notes,
		contact.CreationDate,
		contact.IsDeleted,
	).Scan(&contactId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate primary key error
			return model.Contact{}, model.ErrContactExist
		default:
			return model.Contact{}, model.ErrServiceError
		}
	} else if err != nil {
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

func getCreateContactQuery(ownerId uint64) string {
	return fmt.Sprintf(`
	INSERT INTO %s ("owner_id", "employee_id", "notes", "creation_date", "is_deleted")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING "id";`, getShardName(ownerId))
}

func (c *coreRepoImpl) UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error) {
	if e, err := c.Exec(ctx, getUpdateContactQuery(ownerId),
		contactId,
		upd.Notes,
	); err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Contact{}, model.ErrContactNotExists
	}

	return c.GetContactById(ctx, ownerId, contactId)
}

func getUpdateContactQuery(ownerId uint64) string {
	return fmt.Sprintf(`
		UPDATE %s
		SET "notes" = $2
		WHERE "id" = $1 AND (NOT "is_deleted");`, getShardName(ownerId))
}

func (c *coreRepoImpl) DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error {
	if e, err := c.Exec(ctx, getDeleteContactQuery(ownerId),
		contactId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrContactNotExists
	} else {
		return nil
	}
}

func getDeleteContactQuery(ownerId uint64) string {
	return fmt.Sprintf(`
		UPDATE %s
		SET "is_deleted" = true
		WHERE "id" = $1 AND (NOT "is_deleted");`, getShardName(ownerId))
}

func (c *coreRepoImpl) GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, error) {
	rows, err := c.Query(ctx, getGetContactsQuery(ownerId),
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
			&contact.Empl.Id,
			&contact.Empl.CompanyId,
			&contact.Empl.FirstName,
			&contact.Empl.SecondName,
			&contact.Empl.Email,
			&contact.Empl.JobTitle,
			&contact.Empl.Department,
			&contact.Empl.CreationDate,
			&contact.Empl.IsDeleted,
		)

		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func getGetContactsQuery(ownerId uint64) string {
	shardName := getShardName(ownerId)
	return fmt.Sprintf(`
		SELECT * FROM %s
		INNER JOIN "employees" ON "employee_id" = "employees"."id"
		WHERE "owner_id" = $1 AND (NOT "%s"."is_deleted")
		LIMIT $2 OFFSET $3;`, shardName, shardName)
}

func (c *coreRepoImpl) GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (model.Contact, error) {
	row := c.QueryRow(ctx, getGetContactByIdQuery(ownerId),
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
		&contact.Empl.Id,
		&contact.Empl.CompanyId,
		&contact.Empl.FirstName,
		&contact.Empl.SecondName,
		&contact.Empl.Email,
		&contact.Empl.JobTitle,
		&contact.Empl.Department,
		&contact.Empl.CreationDate,
		&contact.Empl.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Contact{}, model.ErrContactNotExists
	} else if err != nil {
		return model.Contact{}, errors.Join(model.ErrDatabaseError, err)
	}

	return contact, nil
}

func getGetContactByIdQuery(ownerId uint64) string {
	shardName := getShardName(ownerId)
	return fmt.Sprintf(`
		SELECT * FROM %s
		INNER JOIN "employees" ON "employee_id" = "employees"."id"
		WHERE "id" = $1 AND (NOT "%s"."is_deleted");`, shardName, shardName)
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
