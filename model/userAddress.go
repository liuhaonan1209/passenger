package model

import "gorm.io/gorm"

type UserAddress struct { //用户地址表
	gorm.Model
	UserId        uint   `gorm:"type:bigint;not null;comment:关联用户ID"`
	AddressType   int    `gorm:"type:tinyint;not null;comment:地址类型：1 - 家庭，2 - 公司，3 - 常用，4 - 历史"`
	Province      string `gorm:"type:varchar(20);not null;comment:省份"`
	City          string `gorm:"type:varchar(20);not null;comment:城市"`
	District      string `gorm:"type:varchar(20);not null;comment:区县"`
	DetailAddress string `gorm:"type:varchar(255);not null;comment:详细地址"`
	//Latitude      float64 `gorm:"type:decimal(10,6);comment:纬度（GPS 定位）"`
	//Longitude     float64 `gorm:"type:decimal(10,6);comment:经度（GPS 定位）"`
	IsDefault int `gorm:"type:tinyint;not null;default:0;comment:是否默认地址：0 - 否，1 - 是"`
}
