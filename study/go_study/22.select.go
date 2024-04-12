package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义channel
var moneyChan1 = make(chan int) //声明并初始化一个长度为0的信道
var nameChan1 = make(chan string)
var doneChan = make(chan int)

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	start := time.Now()

	wait.Add(3)
	go send("张三", 3, &wait)
	go send("李四", 5, &wait)
	go send("王五", 2, &wait)

	go func() {
		wait.Wait()
		close(moneyChan1)
		close(nameChan1)
		close(doneChan)
	}()

	var moneyList []int
	var nameList []string

	//接收通道中数据的另一种方式，与遍历通道的方式的相比可以在一个循环中接收多个通道，
	for {
		select {
		case name := <-nameChan1: //nameChan1中有数据
			nameList = append(nameList, name)
		case money := <-moneyChan1: //moneyChan1中有数据，通道关闭后不能发送数据但还可以接收通道中的剩余数据，如果没有数据接收操作将接收通道类型的0值
			moneyList = append(moneyList, money)
		case <-doneChan: //doneChan被关闭
			fmt.Println("结束")
			fmt.Println(moneyList)
			fmt.Println(nameList)
			fmt.Println("购物完成", time.Since(start))
			return
		}
	}

}
