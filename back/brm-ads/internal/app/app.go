package app

import (
	"brm-ads/internal/adapters/grpccore"
	"brm-ads/internal/adapters/grpcleads"
	"brm-ads/internal/model"
	"brm-ads/internal/repo"
	"brm-ads/pkg/logger"
	"context"
	"errors"
	"time"
)

type appImpl struct {
	repo  repo.AdRepo
	core  grpccore.CoreClient
	leads grpcleads.LeadsClient

	logs logger.Logger
}

func (a *appImpl) GetAdById(ctx context.Context, id uint64) (model.Ad, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"ad_id":  id,
			"Method": "GetAdById",
		}, err)
	}()

	ad, err := a.repo.GetAdById(ctx, id)
	return ad, err
}

func (a *appImpl) GetAdsList(ctx context.Context, params model.AdsListParams) ([]model.Ad, uint, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "GetAdsList",
		}, err)
	}()
	if params.Filter != nil && params.Filter.ByCompany {
		if _, err = a.core.GetCompany(ctx, params.Filter.CompanyId); err != nil {
			return []model.Ad{}, 0, err
		}
	}

	ads, amount, err := a.repo.GetAdsList(ctx, params)
	return ads, amount, err
}

func (a *appImpl) CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad model.Ad) (model.Ad, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "CreateAd",
		}, err)
	}()

	// setting ad fields
	ad.CompanyId = companyId
	ad.CreatedBy = employeeId
	ad.Responsible = employeeId
	ad.CreationDate = time.Now().UTC()
	ad.IsDeleted = false

	ad, err = a.repo.CreateAd(ctx, ad)
	return ad, err
}

func (a *appImpl) UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd model.UpdateAd) (model.Ad, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "UpdateAd",
		}, err)
	}()

	ad, err := a.repo.GetAdById(ctx, adId)
	if err != nil {
		return model.Ad{}, err
	}
	if ad.Responsible != employeeId {
		return model.Ad{}, model.ErrAuthorization
	}
	var newResponsibleCompanyId uint64
	if newResponsibleCompanyId, _, err = a.core.GetEmployeeById(ctx, companyId, employeeId, upd.Responsible); err != nil {
		return model.Ad{}, err
	} else if newResponsibleCompanyId != ad.CompanyId {
		return model.Ad{}, model.ErrAuthorization
	}

	ad, err = a.repo.UpdateAd(ctx, adId, upd)
	return ad, err
}

func (a *appImpl) DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "DeleteAd",
		}, err)
	}()

	ad, err := a.repo.GetAdById(ctx, adId)
	if err != nil {
		return err
	}
	if ad.Responsible != employeeId {
		return model.ErrAuthorization
	}

	err = a.repo.DeleteAd(ctx, adId)
	return err
}

func (a *appImpl) CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (model.Response, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"ad_id":       adId,
			"Method":      "CreateResponse",
		}, err)
	}()

	ad, err := a.repo.GetAdById(ctx, adId)
	if err != nil {
		return model.Response{}, err
	}
	if ad.CompanyId == companyId {
		return model.Response{}, model.ErrSameCompany
	}

	resp, err := a.repo.CreateResponse(ctx, model.Response{
		CompanyId:    companyId,
		EmployeeId:   employeeId,
		AdId:         adId,
		CreationDate: time.Now().UTC(),
	})

	if a.leads.CreateLead(ctx, adId, companyId, employeeId) != nil {
		return model.Response{}, model.ErrLeadsServiceError
	}

	return resp, err
}

func (a *appImpl) GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]model.Response, uint, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"company_id":  companyId,
			"employee_id": employeeId,
			"Method":      "GetResponses",
		}, err)
	}()

	resps, amount, err := a.repo.GetResponses(ctx, companyId, limit, offset)
	return resps, amount, err
}

func (a *appImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	return a.repo.GetIndustries(ctx)
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrCoreError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
