package main

import (
	"fmt"

	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli"
)

func main() {
	basic.Init()
	conn := rediscli.GetRedis()
	fmt.Println(conn)
}
