package model

import "gorm.io/gorm"

type Coupon struct { //优惠券基础表
	gorm.Model
	CouponName   string  `gorm:"type:varchar(50);not null;comment:优惠券名称（如：新人立减 20 元）"`
	CouponType   int     `gorm:"type:tinyint;not null;comment:优惠券类型：1 - 新人专享，2 - 每日签到，3 - 邀请好友"`
	Denomination float64 `gorm:"type:decimal(10,2);not null;comment:优惠券面额"`
	UseThreshold float64 `gorm:"type:decimal(10,2);not null;default:0;comment:使用门槛（满 X 元可用）"`
	StartTime    string  `gorm:"type:datetime;not null;comment:有效期开始时间"`
	EndTime      string  `gorm:"type:datetime;not null;comment:有效期结束时间"`
	CarTypeIds   string  `gorm:"type:varchar(100);not null;comment:适用车型 ID（逗号分隔，如：1,2）"`
	IsEnable     int     `gorm:"type:tinyint;not null;default:1;comment:是否启用：1 - 是，0 - 否"`
}
