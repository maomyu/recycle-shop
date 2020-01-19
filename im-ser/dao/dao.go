package dao

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/streadway/amqp"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/im-ser/service/model"
)

type WsDao struct {
}

func (w WsDao) UpdateOnlineUser(online string, conn redis.Conn, id string) {
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}
	if _, err := conn.Do("MULTI"); err != nil {
		fmt.Errorf("updateosnlineuser [开启事务]: [%w]", err)
	}
	if _, err := conn.Do("SADD", online, id); err != nil {
		fmt.Errorf("updateonlineuser [添加在线用户]: [%w]", err)
	}
	if _, err := conn.Do("EXEC"); err != nil {
		fmt.Errorf("uperrdateonlineuser [事务执行]: [%w]", err)
	}
}

// 根据用户id获取该用户的状态,用户不存在无状态的情况
func (w WsDao) GetOnlineStatus(online string, conn redis.Conn, id string) int {
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}
	status := 0
	if reply, _ := redis.String(conn.Do("HGET", online, id)); len(reply) > 0 {
		status, _ = strconv.Atoi(reply)
	}

	return status
}

// 判断用户是否处于在线
func (w WsDao) IsOnline(online string, conn redis.Conn, id string) bool {
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}

	if ok, _ := redis.Int(conn.Do("SISMEMBER", online, id)); ok == 1 {
		return true
	}
	return false
}

// 保存历史消息通过消息队列
func SaveHistoryMessage(delivery amqp.Delivery) {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	var message model.Message
	json.Unmarshal(delivery.Body, &message)

	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("updateonlineuser [选择数据库]: [%w]", err)
	}

	score := time.Now().Unix()
	member := string(delivery.Body)
	key := "history:message:" + message.ReceiveID

	// 开始保存
	_, err := conn.Do("ZADD", key, score, member)
	if err == nil {
		delivery.Acknowledger.Ack(delivery.DeliveryTag, true)
	}
}
