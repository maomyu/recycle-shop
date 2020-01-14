package redispool

import (
	"io"

	"github.com/garyburd/redigo/redis"
)

// 新创建一个连接会话
func NewSession() (*redisConnection, error, *Pool, io.Closer) {
	conn, err := p.Acquire()
	if err != nil {
		return nil, err, p, conn.(*redisConnection)
	}
	fa := conn.(*redisConnection)
	return fa, err, p, conn.(*redisConnection)
}
func (s *redisConnection) Relase(p *Pool, c io.Closer) {
	p.Release(c)
}
func (f *redisConnection) GetID() int32 {
	return f.id
}
func (f *redisConnection) GetConn() *redis.Conn {
	return f.conn
}
