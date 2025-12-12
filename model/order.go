package model

import "gorm.io/gorm"

type Order struct { //订单表
	gorm.Model
	OrderNo       string  `gorm:"type:varchar(255);not null;unique;comment:订单编号"`
	UserId        uint    `gorm:"type:bigint;not null;comment:关联用户ID"`
	CarTypeId     uint    `gorm:"type:bigint;not null;comment:关联车型ID"`
	StartProvince string  `gorm:"varchar(20);not null;comment:出发地省份"`
	StartCity     string  `gorm:"varchar(20);not null;comment:出发地城市"`
	StartDistrict string  `gorm:"varchar(20);not null;comment:出发地区县"`
	StartDetail   string  `gorm:"varchar(255);not null;comment:出发地详细地址"`
	EndProvince   string  `gorm:"varchar(20);not null;comment:目的地省份（修改后更新）"`
	EndCity       string  `gorm:"varchar(20);not null;comment:目的地城市（修改后更新）"`
	EndDistrict   string  `gorm:"varchar(20);not null;comment:目的地区县（修改后更新）"`
	EndDetail     string  `gorm:"varchar(255);not null;comment:目的地详细地址（修改后更新）"`
	OrderStatus   int     `gorm:"type:tinyint(1);not null;default:0;comment:订单状态：0 - 待接单，1 - 已接单，2 - 行程中，3 - 已完成，4 - 已取消，5 - 待支付"`
	EstimatePrice float64 `gorm:"decimal(10,2);not null;comment:预估价格"`
	ActualPrice   float64 `gorm:"decimal(10,2);comment:实际支付金额（行程结束后填充）"`
	DriverId      uint    `gorm:"type:bigint;comment:关联司机ID"`
	DriverName    string  `gorm:"varchar(50);comment:司机姓名（脱敏，如：张 *）"`
	DriverMobile  string  `gorm:"char(11);comment:司机手机号（加密存储）"`
	LicensePlate  string  `gorm:"varchar(20);comment:车牌号"`
	CarPhotoUrl   string  `gorm:"varchar(255);comment:车辆照片URL"`
	//OrderCreateTime  string  `gorm:"type:datetime;not null;comment:下单时间"`
	DriverAcceptTime string  `gorm:"type:datetime;comment:司机接单时间"`
	TripStartTime    string  `gorm:"type:datetime;comment:行程开始时间（司机点击 “开始计费”）"`
	TripEndTime      string  `gorm:"type:datetime;comment:行程结束时间（司机点击 “结束计费”）"`
	CancelTime       string  `gorm:"type:datetime;comment:取消订单时间"`
	CancelReason     string  `gorm:"varchar(100);comment:取消原因"`
	CancelFee        float64 `gorm:"decimal(10,2);default:0;comment:取消费用"`
}
