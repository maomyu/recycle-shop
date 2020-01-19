package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/basic/mq"
	"github.com/yuwe1/recycle-shop/im-ser/controller/ws"
	"github.com/yuwe1/recycle-shop/im-ser/dao"
)

func route() {
	// http.HandleFunc("/ws", ws.ServerWs)

}
func main() {
	basic.Init()
	client := mq.GetRabbitMQ()
	fmt.Println(client)
	go client.ConsumeFromQueue("savehistory:234567890", "savehistory", dao.SaveHistoryMessage)
	r := gin.Default()
	r.GET("/ws", ws.ServerWs)
	r.POST("/send", ws.SendMessage)
	http.ListenAndServe("0.0.0.0:8080", r)
}
