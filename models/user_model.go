package models

type User struct {
	UserName string
	PassWord string `json:"password_digest"`
	BaseModel
}
