package models

type User struct {
	UserName string `gorm:"column:UserName"`
	PassWord string `json:"password_digest" gorm:"column:password_digest"`
	BaseModel
}

type UserJson struct {
	UserName string `gorm:"column:UserName"`
	BaseModel
}

func (u User) Json() UserJson {
	return UserJson{
		BaseModel: u.BaseModel,
		UserName:  u.UserName,
	}
}
