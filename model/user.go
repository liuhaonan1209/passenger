package model

import (
	"gorm.io/gorm"
)

type User struct { //用户表
	gorm.Model
	Username   string `gorm:"type:varchar(50);not null;comment:用户昵称"`
	Mobile     string `gorm:"type:char(11);unique;not null;comment:手机号"`
	AvatarUrl  string `gorm:"type:varchar(255);comment:头像URL"`
	Gender     int    `gorm:"type:tinyint(1);default:0;comment:性别：0 - 未知，1 - 男，2 - 女"`
	UserStatus int    `gorm:"type:tinyint(1);default:1;comment:用户状态：1 - 正常，2 - 禁用"`
}
