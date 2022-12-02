package user_helper

import (
	"bbdate/internal/bbdate/domain/model"
	"bbdate/internal/bbdate/infrastructure/http/session"

	"github.com/gin-gonic/gin"
)

const (
	KeyXRequestID string = "key-x-request-id"
	KeyUseryID    string = "key-user-id"
	KeyUserName   string = "key-user-name"
)

func GetUserSessionValueFromGinContext(c *gin.Context) *model.UserContextData {
	var (
		xrid     string
		userID   string
		userName string
	)
	i, ok := c.Get(KeyXRequestID)
	if ok {
		xrid = i.(string)
	}
	i, ok = c.Get(KeyUseryID)
	if ok {
		userID = i.(string)
	}
	i, ok = c.Get(KeyUserName)
	if ok {
		userName = i.(string)
	}
	return model.NewUserContextData(
		xrid,
		userID,
		userName,
	)
}

func BindSessionData(c *gin.Context, ts *session.TmxSession) {
	c.Set(KeyUseryID, ts.UserID)
	c.Set(KeyUserName, ts.UserName)
}
