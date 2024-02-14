package authrepo

import (
	"auth/internal/model"
	"context"
)

type AuthRepo interface {
	SetTokens(ctx context.Context, tokens model.TokensPair) error
	GetTokens(ctx context.Context, access string) (model.TokensPair, error)
	DeleteTokens(ctx context.Context, access string) error
}
