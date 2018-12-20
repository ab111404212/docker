package main

import (
	"deploy/cache"
	"deploy/model"
	"deploy/queue"
	"deploy/util"
	"math/rand"
	"time"
)

const (
	SERVER_ADDR = "127.0.0.1:6379"
)

var end chan interface{}

func main() {

	util.MarshalTest()

	cache.Init()
	AddTask()

	queue.ProcessQueue()

	<-end
}

func AddTask() {
	go func() {
		tick := time.Tick(3 * time.Second)
		for {
			select {
			case <-tick:
				n := rand.Intn(1000)
				v := model.Value{A: n}
				queue.Push(v)
			}
		}
	}()

	// go func() {
	// 	tick := time.Tick(10 * time.Millisecond)
	// 	for {
	// 		select {
	// 		case <-tick:
	// 			n := rand.Intn(1000000)
	// 			v := queue.Value{A: n}
	// 			queue.Push(v)
	// 		}
	// 	}
	// }()

}
