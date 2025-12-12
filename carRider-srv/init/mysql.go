package init

import (
	"carRider-srv/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() {
	mysqlConfig := global.AppCfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password,
		mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database,
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	global.DB.AutoMigrate()
	fmt.Println("迁移成功")
}
