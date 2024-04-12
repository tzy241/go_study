package main

import (
	"fmt"
	"sync"
	"time"
)

var sum int
var wait sync.WaitGroup
var lock sync.Mutex

func addNum() {
	lock.Lock()
	for i := 0; i < 10000; i++ {
		sum++
	}
	lock.Unlock()
	wait.Done()
}

func sub() {
	lock.Lock()
	for i := 0; i < 10000; i++ {
		sum--
	}
	lock.Unlock()
	wait.Done()
}

func main() {
	start := time.Now()
	wait.Add(2)
	go addNum()
	go sub()
	wait.Wait()

	fmt.Println("执行结束", time.Since(start))
	fmt.Println(sum)
}
