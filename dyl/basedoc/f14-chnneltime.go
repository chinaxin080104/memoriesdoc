// 协成超时处理
package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("event-start")
	//协成睡了2秒
	time.Sleep(2 * time.Second)
	fmt.Println("event-end")
	close(done)
}

func main() {
	go event()

	select {
	case <-done:
		fmt.Println("协程执行完毕")

	case <-time.After(3 * time.Second):
		//如果超过3秒，提示超时
		fmt.Println("超时")
		return
	}
}
