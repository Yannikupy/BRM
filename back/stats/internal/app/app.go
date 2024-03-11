package app

import (
	"context"
	"errors"
	"stats/internal/model"
	"stats/internal/repo"
	"stats/pkg/logger"
)

type appImpl struct {
	repo repo.Repo

	logs logger.Logger
}

func (a *appImpl) GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": id,
			"Method":     "GetCompanyMainPageStats",
		}, err)
	}()

	data, err := a.repo.GetCompanyMainPageStats(ctx, id)
	return data, err
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrCoreDatabase) || errors.Is(err, model.ErrLeadsDatabase) || errors.Is(err, model.ErrAdsDatabase) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
