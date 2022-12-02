package helper

import (
	"bbdate/pkg/xrid"

	"github.com/gin-gonic/gin"
)

func GenerateXRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		xrid := xrid.NewXRID()
		BindXRequestID(c, xrid)
		c.Next()
	}
}
