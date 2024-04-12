package main

import (
	"fmt"
	"time"
)

var done = make(chan int)

func event() {
	fmt.Println("event 执行开始")
	time.Sleep(2 * time.Second)
	fmt.Println("event 执行结束")
	close(done)
}

func main() {
	go event()

	select {
	case <-done: //done被关闭
		fmt.Println("协程执行完毕")
	case <-time.After(1 * time.Second): //1秒后执行该分支
		fmt.Println("超时")
	}

}
