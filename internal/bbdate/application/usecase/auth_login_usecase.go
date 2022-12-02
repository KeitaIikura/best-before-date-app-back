package bbdate_usecase

import (
	"bbdate/internal/bbdate/domain/service"
	"bbdate/internal/bbdate/infrastructure/http/session"
	"bbdate/pkg/apperr"
	"bbdate/pkg/logging"
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// data interdace
type (
	AuthLoginUCInput struct {
		ID       string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 変更時はログイン状態の判定の処理の出力も修正すること
	AuthLoginUCOutput struct {
		UserName string `json:"user_name"`
	}
)

// uc interface/interactor
type (
	AuthLoginUsecase interface {
		Execute(ctx *gin.Context, xrid string, input AuthLoginUCInput) (*AuthLoginUCOutput, error)
	}
	AuthLoginInteractor struct {
		AuthService service.AuthService
	}
)

func NewAuthLoginInteractor(
	authService service.AuthService,
) AuthLoginUsecase {
	return &AuthLoginInteractor{
		AuthService: authService,
	}
}

func (i *AuthLoginInteractor) Execute(ctx *gin.Context, xrid string, input AuthLoginUCInput) (*AuthLoginUCOutput, error) {

	user, err := i.AuthService.Login(ctx, xrid, input.ID, input.Password)
	if err != nil {
		logging.Info(xrid, fmt.Sprintf("MngAuthLoginUsecase.Execute login Error: %v", err.Error()))
		return nil, apperr.NewTmxError(apperr.ErrCodeUnauthorized, err)
	}

	// セッション格納
	tmxSession := session.NewTmxSession(strconv.FormatInt(user.ID, 10), user.UserName)
	s := sessions.DefaultMany(ctx, session.TmxSessionKey)
	binaryTmxSession, err := session.MarshalTmxSession(*tmxSession)
	if err != nil {
		logging.Error(xrid, "failed to create session. err: %v", err.Error())
		return nil, apperr.NewTmxError(apperr.ErrCodeUnauthorized, err)
	}
	err = session.SaveToTmxStore(s, binaryTmxSession)
	if err != nil {
		logging.Error(xrid, "failed to save session. err: %v", err.Error())
		return nil, apperr.NewTmxError(apperr.ErrCodeUnauthorized, err)
	}

	return &AuthLoginUCOutput{
		UserName: user.UserName,
	}, nil
}
