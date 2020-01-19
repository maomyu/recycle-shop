package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuwe1/recycle-shop/basic"
	"github.com/yuwe1/recycle-shop/im-ser/controller/ws"
)

func route() {
	// http.HandleFunc("/ws", ws.ServerWs)

}
func main() {
	basic.Init()
	route()
	r := gin.Default()
	r.GET("/ws", ws.ServerWs)
	r.POST("/send", ws.SendMessage)
	http.ListenAndServe("0.0.0.0:8080", r)
}
