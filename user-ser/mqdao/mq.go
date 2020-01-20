package mqdao

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/streadway/amqp"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
	"github.com/yuwe1/recycle-shop/basic/common"
)

func UpdateCreditscore(delivery amqp.Delivery) {
	var creditsocre common.CrediteScore
	json.Unmarshal(delivery.Body, &creditsocre)
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("UpdateCreditscore [选择数据库]: [%w]", err)
	}
	cr := "user:creditscore:"
	// 获取原来的值

	if ok, err := redis.Bool(conn.Do("HINCRBY", cr, creditsocre.ID, creditsocre.Score)); !ok {
		fmt.Errorf("UpdateCreditscore [修改信用值]: [%w]", err)
	} else {
		delivery.Acknowledger.Ack(delivery.DeliveryTag, true)
	}
}
