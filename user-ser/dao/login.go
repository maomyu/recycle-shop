package dao

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
)

func (d UserDao) GetFollowNum(id string) int {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("GetFollowNum [选择数据库]: [%w]", err)
	}
	follow := "user:follow:" + id
	reply, _ := redis.Strings(conn.Do("SMEMBERS", follow))
	return len(reply)
}

// 获得粉丝的数量
func (d UserDao) GetFansNum(id string) int {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("GetFansNum [选择数据库]: [%w]", err)
	}
	fans := "user:fans:" + id
	reply, _ := redis.Strings(conn.Do("SMEMBERS", fans))
	return len(reply)
}

// 获取用户的信用值
func (d UserDao) GetCreditScore(id string) int {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("GetFansNum [选择数据库]: [%w]", err)
	}
	credit := "user:creditscore:"
	reply, _ := redis.Int(conn.Do("HGET", credit, id))
	return reply
}

// 设置用户的在线状态
func (d UserDao) SetOnlineStatus(id string, status int) {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	conn := f.GetConn()
	online := "online:status:"
	conn.Do("HMSET", online, id, status)
}
