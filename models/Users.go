package models

import "gorm.io/gorm"

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
