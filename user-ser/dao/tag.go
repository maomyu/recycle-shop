package dao

import (
	"fmt"

	"github.com/yuwe1/recycle-shop/basic/client/rediscli/redispool"
)

func (t UserDao) UpdateUserTags(ids []string, id string) bool {
	f, _, p, c := redispool.NewSession()
	defer func() {
		if f.GetConn() == nil {
			f.Relase(p, c)
		}
	}()
	tagskey := "user:tags:" + id
	conn := f.GetConn()
	if _, err := conn.Do("SELECT", 0); err != nil {
		fmt.Errorf("GetFansNum [选择数据库]: [%w]", err)
	}
	conn.Do("MULTI")
	for _, v := range ids {
		conn.Do("ZADD", tagskey, 100, v)
	}
	conn.Do("EXEC")
	return true
}
