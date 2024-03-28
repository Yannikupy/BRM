package app

import (
	"context"
	"errors"
	"images/internal/model"
	"images/internal/repo"
	"images/pkg/logger"
	"net/http"
)

type appImpl struct {
	r    repo.ImageRepo
	logs logger.Logger
}

func (a *appImpl) AddImage(ctx context.Context, img model.Image) (uint64, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "AddImage",
		}, err)
	}()

	if len(img) > model.ImageMaxSize {
		err = model.ErrImageTooBig
		return 0, err
	}

	if _, ok := model.PermittedImageTypes[http.DetectContentType(img)]; !ok {
		err = model.ErrWrongFormat
		return 0, err
	}

	id, err := a.r.AddImage(ctx, img)
	return id, err
}

func (a *appImpl) GetImage(ctx context.Context, id uint64) (model.Image, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"Method": "GetImage",
		}, err)
	}()

	img, err := a.r.GetImage(ctx, id)
	return img, err
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) || errors.Is(err, model.ErrServiceError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
