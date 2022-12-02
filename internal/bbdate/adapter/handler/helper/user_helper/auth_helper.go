package user_helper

import (
	"bbdate/internal/bbdate/adapter/handler/helper"
	"bbdate/internal/bbdate/infrastructure/http/session"
	"bbdate/pkg/logging"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		xrid := helper.GetXRequestID(c)
		if isAuthRequired(c.Request.Method, c.Request.URL.Path) {
			s := getSession(c, xrid)
			if s == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}
			BindSessionData(c, s)
			c.Next()
		}
	}
}

// 内部メソッド群

func isAuthRequired(method, path string) bool {
	if method == "OPTIONS" {
		return false
	}
	isAuthRequired := true
	switch path {
	case "/auth/login": // TODO: 定数化
		isAuthRequired = false
	}
	return isAuthRequired
}

func getSession(c *gin.Context, xRequestID string) *session.TmxSession {
	// get session data
	s := sessions.DefaultMany(c, session.TmxSessionKey)
	binarySession := session.GetFromTmxStrore(s)
	ts, err := session.UnmarshalTmxSession(binarySession)
	if err != nil {
		logging.Info(xRequestID, "failed to unmarshal session data")
		return nil
	}

	// 有効期限の更新
	err = session.SaveToTmxStore(s, binarySession)
	if err != nil {
		logging.Info(xRequestID, "failed to save session to redis store")
	}
	// 有効期限延長に失敗しただけなのでセッションは返す
	return ts
}
