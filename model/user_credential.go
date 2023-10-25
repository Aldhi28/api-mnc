package model

type UserCredential struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"isActive"`
}
