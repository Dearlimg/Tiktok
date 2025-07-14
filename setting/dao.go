package setting

import (
	"tiktok001/dao/mysql"
	"tiktok001/dao/redis"
)

type database struct {
}

func (db database) Init() {
	mysql.Init()
	redis.InitRedis()
	rabbitmq
}
