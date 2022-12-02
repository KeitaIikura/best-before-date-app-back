package xrid

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewXRID() string {
	rUUID, err := uuid.NewRandom()
	xRequestID := strings.ReplaceAll(time.Now().Format("20060102150405.000000000"), ".", "")
	if err == nil {
		xRequestID = rUUID.String()
	}
	return xRequestID
}
