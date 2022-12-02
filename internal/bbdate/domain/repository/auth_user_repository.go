package repository

import (
	"bbdate/internal/bbdate/domain/model"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, xrid string, id int64) (*model.User, error)
	GetByEmailAddress(ctx context.Context, xrid string, email_address string) (*model.User, error)
	Create(ctx context.Context, xrid string, mu model.User) error
	Update(ctx context.Context, xrid string, mu model.User) error
}
