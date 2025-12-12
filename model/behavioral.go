package model

import "gorm.io/gorm"

type Behavioral struct { //用户行为表
	gorm.Model
	OrderQuantity int     `gorm:"type:int;comment:累计下单数"`
	OrderPrice    float64 `gorm:"type:decimal(10,2);comment:累计消费金额"`
	CouponNumber  int     `gorm:"type:int;comment:优惠券使用数量"`
}
