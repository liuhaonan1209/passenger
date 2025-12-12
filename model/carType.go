package model

import "gorm.io/gorm"

type CarType struct { //车型配置表
	gorm.Model
	CarTypeName     string `gorm:"type:varchar(20);not null;comment:车型名称：快车、顺风车"`
	PriceRule       string `gorm:"type:text;not null;comment:计费规则（JSON 格式：一口价说明、附加费说明）"`
	CarConfig       string `gorm:"type:varchar(255);not null;comment:车型配置（如车辆空间、座椅数）"`
	ServiceStandard string `gorm:"type:varchar(255);not null;comment:服务标准（如司机服务要求）"`
	IsEnable        int    `gorm:"type:tinyint;not null;default:1;comment:是否启用：1 - 是，0 - 否"`
}
