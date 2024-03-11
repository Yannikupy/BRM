package app

import (
	"context"
	"stats/internal/model"
)

type App interface {
	GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error)
}
