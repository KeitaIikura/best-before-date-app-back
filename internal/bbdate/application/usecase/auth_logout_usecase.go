package bbdate_usecase

import (
	"bbdate/internal/bbdate/infrastructure/http/session"
	"bbdate/pkg/apperr"
	"bbdate/pkg/logging"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// uc interface/interactor
type (
	AuthLogoutUsecase interface {
		Execute(ctx *gin.Context) error
	}
	AuthLogoutInteractor struct {
	}
)

func NewAuthLogoutInteractor() AuthLogoutUsecase {
	return &AuthLogoutInteractor{}
}

func (i *AuthLogoutInteractor) Execute(ctx *gin.Context) error {
	s := sessions.DefaultMany(ctx, session.TmxSessionKey)
	err := session.DeleteSession(s)
	if err != nil {
		logging.Error("failed to delete session. err: %v", err.Error())
		return apperr.NewTmxError(apperr.ErrCodeUnauthorized, err)
	}
	return nil
}
