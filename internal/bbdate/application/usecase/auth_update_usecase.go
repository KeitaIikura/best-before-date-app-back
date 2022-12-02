package bbdate_usecase

import (
	"bbdate/internal/bbdate/domain/service"
	"bbdate/pkg/apperr"
	"bbdate/pkg/logging"
	"bbdate/pkg/validator"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	AuthUpdatePasswordUCInput struct {
		UserID      string `json:"management_user_id" binding:"required"`
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
)

type (
	AuthUpdatePasswordUsecase interface {
		Execute(ctx *gin.Context, xrid string, input AuthUpdatePasswordUCInput) error
	}
	AuthUpdatePasswordInteractor struct {
		AuthService service.AuthService
	}
)

func NewAuthUpdatePasswordInteractor(
	authService service.AuthService,
) AuthUpdatePasswordUsecase {
	return &AuthUpdatePasswordInteractor{
		AuthService: authService,
	}
}

func (i *AuthUpdatePasswordInteractor) Execute(ctx *gin.Context, xrid string, input AuthUpdatePasswordUCInput) error {
	userID, err := strconv.ParseInt(input.UserID, 10, 64)
	if err != nil {
		return apperr.NewTmxError(apperr.ErrCodeInvalidRequest, err)
	}

	err = validator.PasswordValidate(input.NewPassword)
	if err != nil {
		logging.Info(xrid, fmt.Sprintf("AuthUpdatePasswordUsecase.validation: %v", err.Error()))
		return apperr.NewTmxError(apperr.ErrCodeInvalidRequest, err)
	}

	// 変更前後のPWが同じ場合はNG
	if input.OldPassword == input.NewPassword {
		err = errors.New("new password must be different from old password")
		logging.Info(xrid, fmt.Sprintf("AuthUpdatePasswordUsecase.validation: %v", err.Error()))
		return apperr.NewTmxError(apperr.ErrCodeInvalidRequest, err)
	}

	err = i.AuthService.UpdatePassword(ctx, xrid, userID, input.OldPassword, input.NewPassword)
	if err != nil {
		return apperr.ConvertToTmxError(xrid, err)
	}
	return nil
}
