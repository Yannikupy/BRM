package repo

import (
	"context"
	"stats/internal/model"
)

type Repo interface {
	GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error)
}
