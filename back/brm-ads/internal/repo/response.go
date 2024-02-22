package repo

import (
	"brm-ads/internal/model"
	"context"
	"errors"
	"fmt"
)

func (a *adRepoImpl) CreateResponse(ctx context.Context, resp model.Response) (model.Response, error) {
	var respId uint64
	if err := a.QueryRow(ctx, getCreateResponseQuery(resp.CompanyId),
		resp.CompanyId,
		resp.EmployeeId,
		resp.AdId,
		resp.CreationDate,
	).Scan(&respId); err != nil {
		return model.Response{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		resp.Id = respId
		return resp, nil
	}
}

func getCreateResponseQuery(companyId uint64) string {
	return fmt.Sprintf(`
		INSERT INTO %s ("company_id", "employee_id", "ad_id", "creation_date")
		VALUES ($1, $2, $3, $4)
		RETURNING "id";`, getShardName(companyId))
}

func (a *adRepoImpl) GetResponses(ctx context.Context, companyId uint64, limit uint, offset uint) ([]model.Response, error) {
	rows, err := a.Query(ctx, getGetResponsesQuery(companyId),
		companyId,
		limit,
		offset)
	if err != nil {
		return []model.Response{}, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	responses := make([]model.Response, 0)
	for rows.Next() {
		var resp model.Response
		_ = rows.Scan(
			&resp.Id,
			&resp.CompanyId,
			&resp.EmployeeId,
			&resp.AdId,
			&resp.CreationDate,
		)
		responses = append(responses, resp)
	}
	return responses, nil
}

func getGetResponsesQuery(companyId uint64) string {
	return fmt.Sprintf(`
		SELECT * FROM %s
		WHERE "company_id" = $1
		LIMIT $2 OFFSET $3;
		`, getShardName(companyId))
}

func getShardName(companyId uint64) string {
	switch companyId % 4 {
	case 0:
		return "responses_shard01"
	case 1:
		return "responses_shard02"
	case 2:
		return "responses_shard03"
	case 3:
		return "responses_shard04"
	default:
		// ахуеть как ты вообще попал сюда?
		return ""
	}
}
