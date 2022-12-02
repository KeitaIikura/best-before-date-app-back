package auth_controller

import (
	"bbdate/internal/bbdate/adapter/handler/helper"
	"bbdate/internal/bbdate/adapter/handler/helper/user_helper"
	bbdate_usecase "bbdate/internal/bbdate/application/usecase"
	"bbdate/pkg/apperr"
	"bbdate/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IAuthController interface {
		Login(c *gin.Context)
		Check(c *gin.Context)
		Logout(c *gin.Context)
		UpdatePassword(c *gin.Context)
	}

	AuthController struct {
		LoginUC          bbdate_usecase.AuthLoginUsecase
		CheckUC          bbdate_usecase.AuthCheckUsecase
		LogoutUC         bbdate_usecase.AuthLogoutUsecase
		UpdatePasswordUC bbdate_usecase.AuthUpdatePasswordUsecase
	}
)

func NewAuthController(
	loginUC bbdate_usecase.AuthLoginUsecase,
	checkUC bbdate_usecase.AuthCheckUsecase,
	logoutUC bbdate_usecase.AuthLogoutUsecase,
	updatePasswordUC bbdate_usecase.AuthUpdatePasswordUsecase,
) IAuthController {
	return &AuthController{
		LoginUC:          loginUC,
		CheckUC:          checkUC,
		LogoutUC:         logoutUC,
		UpdatePasswordUC: updatePasswordUC,
	}
}

func (a *AuthController) Login(c *gin.Context) {
	xrid := helper.GetXRequestID(c)
	logging.Info(xrid, "AuthController.Login()")

	// requestのbind
	input := &bbdate_usecase.AuthLoginUCInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		// invalid parameter
		res := apperr.NewTmxError(apperr.ErrCodeInvalidRequest, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	// ucの実行
	result, err := a.LoginUC.Execute(c, xrid, *input)
	if err != nil {
		res := apperr.ConvertToTmxError(xrid, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (a *AuthController) Check(c *gin.Context) {
	xrid := helper.GetXRequestID(c)
	logging.Info(xrid, "AuthController.Check()")

	// context_dataの取得
	mcd := user_helper.GetUserSessionValueFromGinContext(c)

	// ucの実行
	result, err := a.CheckUC.Execute(c, xrid, mcd)
	if err != nil {
		res := apperr.ConvertToTmxError(xrid, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (a *AuthController) Logout(c *gin.Context) {
	xrid := helper.GetXRequestID(c)
	logging.Info(xrid, "AuthController.Logout()")

	err := a.LogoutUC.Execute(c)
	if err != nil {
		res := apperr.ConvertToTmxError(xrid, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	c.Status(http.StatusOK)
}

func (a *AuthController) UpdatePassword(c *gin.Context) {
	ecd := user_helper.GetUserSessionValueFromGinContext(c)

	logging.Info(ecd.XRequestID, "AuthController.UpdatePassword()")

	input := &bbdate_usecase.AuthUpdatePasswordUCInput{
		UserID: ecd.UserID,
	}
	if err := c.ShouldBindJSON(input); err != nil {
		logging.Info(ecd.XRequestID, err)
		res := apperr.NewTmxError(apperr.ErrCodeInvalidRequest, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	err := a.UpdatePasswordUC.Execute(c, ecd.XRequestID, *input)
	if err != nil {
		res := apperr.ConvertToTmxError(ecd.XRequestID, err)
		c.JSON(res.StatusCode, res.ResponseBody)
		return
	}

	c.Status(http.StatusOK)
}
