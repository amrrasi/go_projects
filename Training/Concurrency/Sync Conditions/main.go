package main

import (
	"sync"
	"time"
)

var usrList = []int{}
var ready = false

func main() {

	Streaming()

}

func Streaming() {
	condition := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 1000; i++ {
		go NewRequest(i, condition)
	}
}

func NewRequest(userId int, condition *sync.Cond) {
	checking(userId, condition)

	condition.L.Lock()
	defer condition.L.Unlock()

	for !ready {
		condition.Wait()
	}
	println("user", userId, " start streaming")

}
func checking(userId int, condition *sync.Cond) {

	println("UserId ", userId, " Waiting to start streaming")
	time.Sleep(time.Millisecond * 300)
	condition.L.Lock()
	defer condition.L.Unlock()

	usrList = append(usrList, userId)
	if len(usrList) == 55 {
		ready = true
		condition.Broadcast()
		println("Streaming Started", userId)
	}

}
