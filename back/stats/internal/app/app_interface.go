package app

import (
	"context"
	"stats/internal/model"
	"stats/internal/repo"
	"stats/pkg/logger"
)

type App interface {
	GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error)
}

func New(repo repo.Repo, logs logger.Logger) App {
	return &appImpl{
		repo: repo,
		logs: logs,
	}
}
