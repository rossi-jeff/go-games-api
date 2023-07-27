package models

import "base_model"

type User struct {
	UserName string
	PassWord string `json:"password_digest"`
	BaseModel
}