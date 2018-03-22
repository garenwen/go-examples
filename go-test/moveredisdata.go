package main

import (
	"time"
	"github.com/garyburd/redigo/redis"
)

type Cache struct {
	redisPool *redis.Pool
	prefix    string
}

var gCache Cache

func (self *Cache) Init(IP, Port, Password string) bool {
	self.redisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", IP+":"+Port)
			if err != nil {
				return nil, err
			}
			if len(Password) > 0 {
				if _, err := c.Do("AUTH", Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return true
}

func (self *Cache) GetConn() redis.Conn {
	return self.redisPool.Get()
}

func MoveRedisData() {

}
