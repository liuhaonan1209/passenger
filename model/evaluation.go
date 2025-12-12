package model

import "gorm.io/gorm"

type Evaluation struct { //评价
	gorm.Model
	OrderNo       string `gorm:"type:varchar(255);not null;unique;comment:订单编号"`
	UserId        uint   `gorm:"type:bigint;not null;comment:关联用户ID"`
	DriverId      uint   `gorm:"type:bigint;not null;comment:关联司机ID"`
	ServiceScore  int    `gorm:"type:tinyint;not null;comment:服务态度评分（1-5 星）"`
	DrivingScore  int    `gorm:"type:tinyint;not null;comment:驾驶技术评分（1-5 星）"`
	CarCleanScore int    `gorm:"type:tinyint;not null;comment:车辆整洁度评分（1-5 星）"`
	Content       string `gorm:"type:varchar(100);comment:文字评价（最多 100 字）"`
	IsAnonymous   int    `gorm:"type:tinyint;not null;default:0;comment:是否匿名：0 - 否，1 - 是"`
}
