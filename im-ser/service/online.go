package service

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/gorilla/websocket"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/basic/mq"
	"github.com/yuwe1/recycle-shop/im-ser/dao"
	"github.com/yuwe1/recycle-shop/im-ser/service/model"
)

var m map[string]*websocket.Conn

func init() {
	m = make(map[string]*websocket.Conn, math.MaxInt64)
}

type UserWs struct {
	ID string
	//在线状态
	Status int
}

func (u *UserWs) UpdateOnlineUser(wc *websocket.Conn) (bool, error) {
	f, err, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()

	online := "online:"
	conn := f.GetConn()
	wsdao := dao.WsDao{}
	// 更新在线用户
	wsdao.UpdateOnlineUser(online, conn, u.ID)
	// 获得用户的状态
	status := wsdao.GetOnlineStatus(online+"status:", conn, u.ID)
	u.Status = status
	m[u.ID] = wc
	return true, err
}

// 定义一个reader监听客户端发送的消息
func (u UserWs) Reader(conn *websocket.Conn) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.Sugar.Errorf("reader: [%w]", err)
			return fmt.Errorf("reader: [%w]", err)
		}
		//定义一个消息
		ms := new(model.Message)
		// 解析消息
		json.Unmarshal(p, ms)
		if err := conn.WriteMessage(messageType, p); err != nil {
			logger.Sugar.Errorf("reader: [%w]", err)
			return fmt.Errorf("reader: [%w]", err)
		}
	}
}
func (u *UserWs) SendMessage(msg model.Message) bool {
	// 根据id找到ws连接
	// ws := m[u.ID]
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	online := "online:"
	conn := f.GetConn()
	wsdao := dao.WsDao{}

	// 获取对方的在线状态值
	status := wsdao.GetOnlineStatus(online+"status:", conn, msg.ReceiveID)
	u.Status = status
	msg.Status = u.Status
	// 不在线， 将消息存储在redis中，的有序集合中，并时间戳当作分值
	if status != 0 {
		m[msg.ReceiveID].WriteJSON(&msg)
	} else {
		// 不在线，消息队列保存发布条消息，由订阅方进行redis存储
		client := mq.GetRabbitMQ()
		byt, _ := json.Marshal(&msg)
		client.PublishOnQueue(byt, "im-ser", "savehistory", "savehistory:"+msg.ReceiveID)
	}
	return true
}
