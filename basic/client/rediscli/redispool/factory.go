package redispool

import (
	"io"

	"github.com/garyburd/redigo/redis"
)

type factory struct {
	id   int32
	conn *redis.Conn //连接的标志
}

// 新创建一个连接会话
func NewSession() (*factory, error, *Pool, io.Closer) {
	conn, err := p.Acquire()
	if err != nil {
		return nil, err, p, conn.(*redisConnection)
	}
	fa := new(factory)
	fa.id = conn.(*redisConnection).id
	fa.conn = conn.(*redisConnection).conn
	return fa, err, p, conn.(*redisConnection)
}
func (s *factory) Relase(p *Pool, c io.Closer) {
	p.Release(c)
}
func (f *factory) GetID() int32 {
	return f.id
}
func (f *factory) GetConn() *redis.Conn {
	return f.conn
}
