package init

import (
	"carRider-srv/global"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Redis() {
	redisConfig := global.AppCfg.Redis
	Addr := fmt.Sprintf("%s:%d",
		redisConfig.Host, redisConfig.Port,
	)
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
	fmt.Println("redis 连接成功")
}
