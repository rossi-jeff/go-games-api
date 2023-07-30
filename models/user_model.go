package models

type User struct {
	UserName string `gorm:"column:UserName"`
	PassWord string `json:"password_digest" gorm:"column:password_digest"`
	BaseModel
}
