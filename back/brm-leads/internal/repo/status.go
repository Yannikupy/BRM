package repo

import (
	"brm-leads/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

const (
	getStatusesQuery = `
		SELECT * FROM "status";`

	getStatusByIdQuery = `
		SELECT "name" FROM "status"
		WHERE "id" = $1;`
)

func (l *leadRepoImpl) GetStatuses(ctx context.Context) (map[string]uint64, error) {
	rows, err := l.Query(ctx, getStatusesQuery)
	if err != nil {
		return map[string]uint64{}, model.ErrDatabaseError
	}
	defer rows.Close()

	statuses := make(map[string]uint64)
	for rows.Next() {
		var id uint64
		var status string
		_ = rows.Scan(&id, &status)
		statuses[status] = id
	}
	return statuses, nil
}

func (l *leadRepoImpl) GetStatusById(ctx context.Context, id uint64) (string, error) {
	row := l.QueryRow(ctx, getStatusByIdQuery, id)
	var status string
	if err := row.Scan(&status); errors.Is(err, pgx.ErrNoRows) {
		return "", model.ErrStatusNotExists
	} else if err != nil {
		return "", errors.Join(model.ErrDatabaseError, err)
	} else {
		return status, nil
	}
}
