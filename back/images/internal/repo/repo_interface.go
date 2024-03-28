package repo

import (
	"context"
	"images/internal/model"
)

type ImageRepo interface {
	AddImage(ctx context.Context, img model.Image) (uint64, error)
	GetImage(ctx context.Context, id uint64) (model.Image, error)
}
