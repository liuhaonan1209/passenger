package model

import "gorm.io/gorm"

type PayRecord struct { //支付记录表
	gorm.Model
	OrderNo   string  `gorm:"type:varchar(255);not null;unique;comment:订单编号"`
	UserId    uint    `gorm:"type:bigint;not null;comment:关联用户ID"`
	PayAmount float64 `gorm:"type:decimal(10,2);not null;comment:支付金额"`
	PayType   int     `gorm:"type:tinyint;not null;comment:支付方式：1 - 微信支付"`
	TradeNo   string  `gorm:"type:varchar(64);not null;unique;comment:第三方交易单号(微信支付单号)"`
	PayStatus int     `gorm:"type:tinyint;not null;default:0;comment:支付状态：0 - 待支付，1 - 支付成功，2 - 支付失败，3 - 退款中，4 - 退款成功"`
	PayTime   string  `gorm:"type:datetime;comment:支付成功时间"`
}
