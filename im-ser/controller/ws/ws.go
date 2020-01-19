package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yuwe1/recycle-shop/basic/common"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/im-ser/service"
	"github.com/yuwe1/recycle-shop/im-ser/service/model"
)

// 返回信息
type Result struct {
	// 错误提示信息
	Tip string `json:"tips"`
}

// 定义一个upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 暂时不需要对客户端进行检查
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocker服务端接口函数上线的入口
func ServerWs(c *gin.Context) {
	// 获得用户的唯一id
	userID := c.Query("id")
	r := c.Request
	w := c.Writer
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Sugar.Error("服务升级失败: [%w]", err)
		result := common.Result{
			Success:   0,
			Errorcode: 1,
			Message: &Result{
				Tip: "服务内部升级错误",
			},
		}
		body, _ := json.Marshal(result)
		w.WriteHeader(500)
		w.Write(body)
	}
	// 将该用户修改成在线用户
	userservice := service.UserWs{
		ID:     userID,
		Status: 1,
	}

	if ok, err := userservice.UpdateOnlineUser(ws); ok && err == nil {
		logger.Sugar.Infof("用户：[%s] 开始等待读取客户端消息", userID)
		userservice.Reader(ws)
	} else {
		logger.Sugar.Error(err)
	}
}

// 用户发送消息
func SendMessage(c *gin.Context) {
	msg := new(model.Message)
	err := c.BindJSON(msg)
	if err != nil {
		fmt.Errorf("[im-ser:controller:ws:sendmessage]:[接收数据出错]:[%w]", err)
	} else {
		if len(msg.ID) > 0 && len(msg.SendID) > 0 && len(msg.ReceiveID) > 0 && len(msg.Content) > 0 && msg.Type > -1 {
			// 定义一个ws客户端
			userws := service.UserWs{
				ID: msg.SendID,
			}
			if userws.SendMessage(*msg) {
				result := common.Result{
					Success:   0,
					Errorcode: 0,
				}
				c.IndentedJSON(200, result)
			} else {
				//
				result := common.Result{
					Success:   0,
					Errorcode: 2,
					Message: &Result{
						Tip: "内部错误",
					},
				}
				c.IndentedJSON(500, result)
			}
		} else {
			result := common.Result{
				Success:   0,
				Errorcode: 1,
				Message: &Result{
					Tip: "无法发送空消息",
				},
			}
			c.IndentedJSON(200, result)
		}

	}
}
