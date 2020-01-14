package main

import (
	"fmt"

	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
)

func main() {
	basic.Init()
	conn := rediscli.GetRedis()
	fmt.Println(conn)
	f, _, p, c := redispool.NewSession()
	f, _, p, c = redispool.NewSession()
	defer f.Relase(p, c)
	fmt.Println(f.GetID())
}
