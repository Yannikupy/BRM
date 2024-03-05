package repo

import (
	"brm-leads/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (l *leadRepoImpl) CreateLead(ctx context.Context, lead model.Lead) (model.Lead, error) {
	var leadId uint64
	if err := l.QueryRow(ctx, createLeadQuery(lead.CompanyId),
		lead.AdId,
		lead.Title,
		lead.Description,
		lead.Price,
		lead.Status,
		lead.Responsible,
		lead.CompanyId,
		lead.ClientCompany,
		lead.ClientEmployee,
		lead.CreationDate,
		lead.IsDeleted,
	).Scan(&leadId); err != nil {
		return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		lead.Id = leadId
		return lead, nil
	}
}

func createLeadQuery(companyId uint64) string {
	return fmt.Sprintf(`
		INSERT INTO %s ("ad_id", "title", "description", "price", "status", "responsible", "company_id", "client_company", "client_employee", "creation_date", "is_deleted")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING "id";`, getShardName(companyId))
}

func (l *leadRepoImpl) GetLeads(ctx context.Context, companyId uint64, filter model.Filter) ([]model.Lead, error) {
	rows, err := l.Query(ctx, getLeadsQuery(companyId),
		companyId,
		filter.ByStatus,
		filter.Status,
		filter.ByResponsible,
		filter.Responsible,
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		return []model.Lead{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	leads := make([]model.Lead, 0)
	for rows.Next() {
		var lead model.Lead
		_ = rows.Scan(
			&lead.Id,
			&lead.AdId,
			&lead.Title,
			&lead.Description,
			&lead.Price,
			&lead.Status,
			&lead.Responsible,
			&lead.CompanyId,
			&lead.ClientCompany,
			&lead.ClientEmployee,
			&lead.CreationDate,
			&lead.IsDeleted,
		)
		leads = append(leads, lead)
	}
	return leads, nil
}

func getLeadsQuery(companyId uint64) string {
	return fmt.Sprintf(`
		SELECT * FROM %s
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "status" = $3)
			AND ((NOT $4) OR "responsible" = $5)
		ORDER BY "creation_date" DESC
		LIMIT $6 OFFSET $7;`, getShardName(companyId))
}

func (l *leadRepoImpl) GetLeadById(ctx context.Context, companyId uint64, id uint64) (model.Lead, error) {
	row := l.QueryRow(ctx, getLeadByIdQuery(companyId), id)
	var lead model.Lead
	if err := row.Scan(
		&lead.Id,
		&lead.AdId,
		&lead.Title,
		&lead.Description,
		&lead.Price,
		&lead.Status,
		&lead.Responsible,
		&lead.CompanyId,
		&lead.ClientCompany,
		&lead.ClientEmployee,
		&lead.CreationDate,
		&lead.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Lead{}, model.ErrLeadNotExists
	} else if err != nil {
		return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return lead, nil
	}
}

func getLeadByIdQuery(companyId uint64) string {
	return fmt.Sprintf(`
		SELECT * FROM %s
		WHERE "id" = $1 AND (NOT "is_deleted");`, getShardName(companyId))
}

func (l *leadRepoImpl) UpdateLead(ctx context.Context, companyId uint64, id uint64, upd model.UpdateLead) (model.Lead, error) {
	if e, err := l.Exec(ctx, updateLeadQuery(companyId),
		id,
		upd.Title,
		upd.Description,
		upd.Price,
		upd.Status,
		upd.Responsible,
	); err != nil {
		return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Lead{}, model.ErrLeadNotExists
	} else {
		return l.GetLeadById(ctx, companyId, id)
	}
}

func updateLeadQuery(companyId uint64) string {
	return fmt.Sprintf(`
		UPDATE %s
		SET "title" = $2,
			"description" = $3,
			"price" = $4,
			"status" = $5,
			"responsible" = $6
		WHERE "id" = $1 AND (NOT "is_deleted");`, getShardName(companyId))
}

func getShardName(companyId uint64) string {
	switch companyId % 4 {
	case 0:
		return "leads_shard01"
	case 1:
		return "leads_shard02"
	case 2:
		return "leads_shard03"
	case 3:
		return "leads_shard04"
	default:
		return ""
	}
}
