package queue

import (
	"log"
	"redis/cache"
	"redis/logger"
	"redis/model"
	"redis/util"
)

const (
	MAX_CHANNLE_SIZE = 20000
	REDIS_QUEUE_NAME = "TestQueue"
)

var queue = make(chan interface{}, MAX_CHANNLE_SIZE)

func Push(v interface{}) {
	log.Println("push...")
	queue <- v
	PushRedis(v)
}

func PopHandler(v interface{}) {
	val := v.(model.Value)
	log.Println("pop....,len=", val.A, len(queue))
	PopRedis()
}

func ProcessQueue() {
	log.Println("ProcessQueue...")
	go func() {
		defer close(queue)
		for {
			select {
			case p := <-queue:
				PopHandler(p)
			}
		}
	}()
}

func PushRedis(v interface{}) {
	c := cache.GetConn()
	defer c.Close()
	b := util.Encoding(v)
	_, err := c.Do("RPUSH", REDIS_QUEUE_NAME, b)
	if err != nil {
		logger.Println(err.Error())
	}
}

func PopRedis() (v interface{}) {
	c := cache.GetConn()
	defer c.Close()
	v, err := c.Do("LPOP", REDIS_QUEUE_NAME)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	b := v.([]byte)
	v = util.Decoding(b)
	// err = json.Unmarshal(b, &v1)
	logger.Println("PopRedis=", v)
	return
}
