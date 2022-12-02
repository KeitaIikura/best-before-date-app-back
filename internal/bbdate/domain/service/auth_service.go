package service

import (
	"bbdate/internal/bbdate/domain/model"
	"bbdate/internal/bbdate/domain/repository"
	"bbdate/pkg/apperr"
	"bbdate/pkg/crypter"
	"bbdate/pkg/logging"
	"context"
	"errors"
	"fmt"
)

type AuthService interface {
	Login(ctx context.Context, xRequestID string, userID string, password string) (*model.AuthUser, error)
	UpdatePassword(ctx context.Context, xrid string, userID int64, oldPw string, newPw string) error
}

type AuthServiceImple struct {
	AuthUserRepo repository.AuthUserRepository
	Crypter      crypter.ICrypter
}

func NewAuthService(
	authUserRepo repository.AuthUserRepository,
	crypter crypter.ICrypter,
) AuthService {
	return &AuthServiceImple{
		AuthUserRepo: authUserRepo,
		Crypter:      crypter,
	}
}

func (s *AuthServiceImple) Login(ctx context.Context, xrid string, userID string, password string) (*model.AuthUser, error) {
	// userID(emailaddress)で検索
	user, err := s.AuthUserRepo.GetByEmailAddress(ctx, xrid, userID)
	if err != nil {
		logging.Error(xrid, fmt.Sprintf("AuthService Login error: %v", err))
		return nil, apperr.NewTmxError(apperr.ErrCodeDBConnection, err)
	}

	// emailaddressが存在しないとき
	if user == nil {
		return nil, apperr.NewTmxError(apperr.ErrCodeUnauthorized, errors.New("not correct auth input"))
	}

	// passwordをハッシュ化して比較
	hashedPassword := s.Crypter.GenerateSha512Hash(password)

	if hashedPassword != user.Password {
		return nil, apperr.NewTmxError(apperr.ErrCodeUnauthorized, errors.New("not correct auth input"))
	}

	// 同じならログイン成功
	return user, nil
}

func (s *AuthServiceImple) UpdatePassword(ctx context.Context, xrid string, userID int64, oldPw string, newPw string) error {
	user, err := s.AuthUserRepo.GetByID(ctx, xrid, userID)
	if err != nil {
		logging.Error(xrid, fmt.Sprintf("AuthService UpdatePassword error: %v", err))
		return apperr.NewTmxError(apperr.ErrCodeDBConnection, err)
	}

	// 現在のパスワードをハッシュ化して比較
	hashedOldPw := s.Crypter.GenerateSha512Hash(oldPw)
	if hashedOldPw != user.Password {
		return apperr.NewTmxError(apperr.TmxErrorCode(apperr.ErrCodeInvalidRequest), errors.New("not correct auth input"))
	}

	// 変更後のパスワードをハッシュ化
	hashedNewPw := s.Crypter.GenerateSha512Hash(newPw)
	user.Password = hashedNewPw

	err = s.AuthUserRepo.Update(ctx, xrid, *user)
	if err != nil {
		return apperr.NewTmxError(apperr.ErrCodeDBConnection, err)
	}

	return nil
}
