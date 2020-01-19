package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/user-ser/controller"
)

func main() {
	basic.Init()
	// client := mq.GetRabbitMQ()
	// fmt.Println(client)
	// go client.ConsumeFromQueue("savehistory:234567890", "savehistory", dao.SaveHistoryMessage)
	// dbpool.GetSession()
	r := gin.Default()
	r.POST("/user/register", controller.Register)
	r.GET("/user/login", controller.Login)
	http.ListenAndServe("0.0.0.0:8081", r)
}
