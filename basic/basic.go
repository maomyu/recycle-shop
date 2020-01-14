package basic

import (
	"github.com/yuwe1/recycle-shop/basic/client/rediscli"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/config"
)

func Init() {
	config.Init()
	rediscli.Init()
	redispool.Init()
}
