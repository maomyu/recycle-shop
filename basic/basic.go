package basic

import (
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/client/dbpool"
	"github.com/yuwe1/recycle-shop/basic/config"
	"github.com/yuwe1/recycle-shop/basic/mq"
)

func Init() {
	config.Init()
	// rediscli.Init()
	redispool.Init()
	mq.Init()
	dbpool.Init()
}
