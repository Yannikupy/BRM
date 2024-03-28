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

func (a *appImpl) UpdateRatingByClosedLead(ctx context.Context, companyId uint64, submit bool) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id": companyId,
			"Method":     "UpdateRatingByClosedLead",
		}, err)
	}()

	currRating, err := a.repo.GetCompanyRating(ctx, companyId)
	if err != nil {
		return err
	}

	var newRating float64
	if submit {
		newRating = increaseRating(currRating)
	} else {
		newRating = decreaseRating(currRating)
	}
	err = a.repo.SetCompanyRating(ctx, companyId, newRating)
	return err
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

func increaseRating(currRating float64) float64 {
	// TODO: implement
	return currRating
}

func decreaseRating(currRating float64) float64 {
	// TODO: implement
	return currRating
}
