package model

type AuthUser struct {
	ID           int64  `json:"id"`
	UserName     string `json:"userName"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

// TODO:関数名からModelの除去
func NewAuthUserModel(userName string, email string, pswd string) *AuthUser {
	return &AuthUser{
		UserName:     userName,
		EmailAddress: email,
		Password:     pswd,
	}
}
