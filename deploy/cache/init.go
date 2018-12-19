package cache

import (
	"log"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

func Init() {
	pool = &redis.Pool{
		MaxIdle:     300,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			//      if _, err := c.Do("AUTH", password); err != nil {
			//        c.Close()
			//        return nil, err
			//      }
			//      if _, err := c.Do("SELECT", db); err != nil {
			//        c.Close()
			//        return nil, err
			//      }
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	testRedis()
}

func GetPool() *redis.Pool {
	return pool
}

func GetConn() redis.Conn {
	return pool.Get()
}

func testRedis() {
	con := GetConn()
	defer con.Close()
	_, err := con.Do("Ping")
	if err != nil {
		log.Fatalln("Redis connect Error")
		return
	}
}
