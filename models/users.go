package models

import (
	"daidai-server/pkg/util"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	// 用户名称
	Name string
	// 用户密码
	Password string
	// 用户手机号
	Phone string
	// 用户邮箱
	Email string
	// 是否锁定
	Locking bool
}

func AddUser(users Users) bool {
	encodePassword := util.EncryptPassword(users.Password)
	users.Password = encodePassword
	DB.Create(users)
	return true
}

func Login(account string, password string) bool {
	users := Users{}
	DB.Where("name = ?", account).First(&users)
	if (users == Users{}) {
		DB.Where("phone = ?", account).First(&users)
	}
	if (users == Users{}) {
		DB.Where("email = ?", account).First(&users)
	}

	if (users == Users{}) {
		return false
	}
	return util.MatchPassword(password, users.Password)
}
