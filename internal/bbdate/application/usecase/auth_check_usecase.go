package bbdate_usecase

import (
	"bbdate/internal/bbdate/domain/model"
	"bbdate/internal/bbdate/domain/service"

	"github.com/gin-gonic/gin"
)

// data interdace
type (
	AuthCheckOutput AuthLoginUCOutput
)

// uc interface/interactor
type (
	AuthCheckUsecase interface {
		Execute(ctx *gin.Context, xrid string, mcd *model.UserContextData) (*AuthCheckOutput, error)
	}
	AuthCheckInteractor struct {
		AuthService service.AuthService
	}
)

func NewAuthCheckInteractor() AuthCheckUsecase {
	return &AuthCheckInteractor{}
}

func (i *AuthCheckInteractor) Execute(ctx *gin.Context, xrid string, mcd *model.UserContextData) (*AuthCheckOutput, error) {

	userName := mcd.UserName

	return &AuthCheckOutput{
		UserName: userName,
	}, nil
}
