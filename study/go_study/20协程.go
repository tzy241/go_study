package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	//wait--
	wait.Done()
}

// var wait int

func main() {
	var wait sync.WaitGroup
	start := time.Now()
	//pay("张三")
	//pay("李四")
	//pay("王五")

	//一起去购物即并发
	//主线程结束，协程函数跟着结束
	//wait = 3

	wait.Add(3)
	go shopping("张三", &wait)
	go shopping("李四", &wait)
	go shopping("王五", &wait)
	wait.Wait() //等待三个协程结束

	//for {
	//	if wait == 0 {
	//		break
	//	}
	//}

	fmt.Println("购物完成", time.Since(start))
}
