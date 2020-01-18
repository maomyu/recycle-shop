package service

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/logger"
	"github.com/yuwe1/recycle-shop/im-ser/service/model"
)

var m map[string]*websocket.Conn

type UserService struct {
	ID string
	//在线状态
	Status int
}

func (u UserService) UpdateOnlineUser(wc *websocket.Conn) (bool, error) {
	f, err, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()

	// 检查redis中是否有该用户
	online := "online:"
	conn := f.GetConn()
	if ok, err := redis.Int(conn.Do("SISMEMBER", online, u.ID)); ok == 1 {
		return false, fmt.Errorf("updateonlineuser: [%w]", err)
	}
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}
	if _, err := conn.Do("MULTI"); err != nil {
		fmt.Errorf("updateonlineuser [开启事务]: [%w]", err)
	}
	if _, err := conn.Do("SADD", online, u.ID); err != nil {
		fmt.Errorf("updateonlineuser [添加在线用户]: [%w]", err)
	}
	// 将用户状态进行存储
	if _, err := conn.Do("HMSET", online+"status:", u.ID, u.Status); err != nil {
		fmt.Errorf("updateonlineuser [添加在线用户的状态]: [%w]", err)
	}
	if _, err := conn.Do("EXEC"); err != nil {
		fmt.Errorf("uperrdateonlineuser [事务执行]: [%w]", err)
	}

	return true, err
}

// 定义一个reader监听客户端发送的消息
func (u UserService) Reader(conn *websocket.Conn) error {
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
