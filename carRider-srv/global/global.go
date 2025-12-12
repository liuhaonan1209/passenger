package global

import (
	"carRider-srv/config"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	AppCfg *config.AppConfig
	Rdb    *redis.Client
	Es     *elasticsearch.Client
)
