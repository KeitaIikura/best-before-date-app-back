package model

type UserContextData struct {
	XRequestID string
	UserID     string
	UserName   string
}

func NewUserContextData(xRequestID string, userID string, userName string) *UserContextData {
	return &UserContextData{
		XRequestID: xRequestID,
		UserID:     userID,
		UserName:   userName,
	}
}
