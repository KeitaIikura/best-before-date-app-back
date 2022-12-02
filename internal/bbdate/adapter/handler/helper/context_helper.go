package helper

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyXRequestID string = "key-x-request-id"
)

func GetXRequestID(c *gin.Context) string {
	var xrid string
	i, ok := c.Get(KeyXRequestID)
	if ok {
		xrid = i.(string)
	}
	return xrid
}

func BindXRequestID(c *gin.Context, xrid string) {
	c.Set(KeyXRequestID, xrid)
}
