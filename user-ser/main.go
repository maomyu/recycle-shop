package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/basic/mq"
	"github.com/yuwe1/recycle-shop/user-ser/controller"
	"github.com/yuwe1/recycle-shop/user-ser/mqdao"
)

func main() {
	basic.Init()
	client := mq.GetRabbitMQ()
	go client.ConsumeFromQueue("creditscore:", "updatecreditscore", mqdao.UpdateCreditscore)
	r := gin.Default()
	r.POST("/user/account", controller.Register)
	r.GET("/user/account/:id", controller.Login)
	// 更新用户的标签
	r.POST("/user/account/tags/:id", controller.UpdateTags)
	http.ListenAndServe("0.0.0.0:8081", r)
}
