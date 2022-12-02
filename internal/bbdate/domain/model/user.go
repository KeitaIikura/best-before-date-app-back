package model

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

func NewUser(name string, email string) *User {
	return &User{
		Name:         name,
		EmailAddress: email,
	}
}
