package model

import "gorm.io/gorm"

type OrderFeeDetail struct { //订单费用明细表
	gorm.Model
	OrderNo        string  `gorm:"type:varchar(255);not null;unique;comment:订单编号"`
	OnePrice       float64 `gorm:"type:decimal(10,2);not null;comment:一口价金额"`
	AdditionalFee  float64 `gorm:"type:decimal(10,2);not null;default:0;comment:附加费（如高速费、停车费）"`
	CouponDeduct   float64 `gorm:"type:decimal(10,2);not null;default:0;comment:优惠券抵扣金额"`
	PrepayDeduct   float64 `gorm:"type:decimal(10,2);not null;default:0;comment:预付金额抵扣"`
	FinalPayAmount float64 `gorm:"type:decimal(10,2);not null;comment:最终实付金额"`
	Mileage        float64 `gorm:"type:decimal(5,2);comment:实际行驶里程（公里）"`
	TripDuration   int     `gorm:"type:int;comment:实际行程时长（分钟）"`
}
