package repository

import (
	"bbdate/internal/bbdate/domain/model"
	"context"
)

type AuthUserRepository interface {
	GetByID(ctx context.Context, xrid string, id int64) (*model.AuthUser, error)
	GetByEmailAddress(ctx context.Context, xrid string, email_address string) (*model.AuthUser, error)
	Create(ctx context.Context, xrid string, mu model.AuthUser) error
	Update(ctx context.Context, xrid string, mu model.AuthUser) error
}
