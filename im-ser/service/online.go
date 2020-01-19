package service

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/logger"
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

func (u UserWs) UpdateOnlineUser(wc *websocket.Conn) (bool, error) {
	f, err, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()

	online := "online:"
	conn := f.GetConn()
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
func (u UserWs) SendMessage(msg model.Message) bool {
	// 根据id找到ws连接
	// ws := m[u.ID]
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	// 根据用户的ID找到status
	online := "online:"
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}
	if reply, _ := redis.String(conn.Do("HGET", online+"status:", u.ID)); len(reply) > 0 {
		status, _ := strconv.Atoi(reply)
		u.Status = status
	} else {
		return false
	}

	// 检查接收方是否处于在线状态,不在线，将把此未读消息放进一个临时的消息队列中消息队列中
	if ok, _ := redis.Int(conn.Do("SISMEMBER", online, msg.ReceiveID)); ok == 1 {
		m[msg.ReceiveID].WriteJSON(&msg)
	} else {
		// 不在线
	}
	return true
}
