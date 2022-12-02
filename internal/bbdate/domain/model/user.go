package model

type User struct {
	ID           int64  `json:"id"`
	UserName     string `json:"user_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

// TODO:関数名からModelの除去
func NewUserModel(userName string, email string, pswd string) *User {
	return &User{
		UserName:     userName,
		EmailAddress: email,
		Password:     pswd,
	}
}
