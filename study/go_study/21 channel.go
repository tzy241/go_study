package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义channel，make(chan 元素类型，缓冲大小)
// 通道有 发送，接受(<-)和关闭(close)操作,当大小为0时，发送后必须接收
// 有缓冲，发送后不一定必须接收
var moneyChan = make(chan int) //声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan <- money

	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	start := time.Now()

	wait.Add(3)
	go pay("张三", 3, &wait)
	go pay("李四", 5, &wait)
	go pay("王五", 2, &wait)

	//并行函数，但是会等待上边三个函数执行完毕后执行close(moneyChan)
	//匿名函数直接调用
	go func() {
		wait.Wait()
		close(moneyChan)
	}()

	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

	var moneyList []int
	//接收通道的一种方式，比较常用，下边是第二种方式，由于通道长度为0，只要通道中有变量，就可以进行接收，当关闭通道后该方式才可以结束循环
	for money := range moneyChan {
		fmt.Println(money)
		moneyList = append(moneyList, money)
	}

	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}
	//无论那种方式的循环，信道中的数据一旦使用后就不存在了
	fmt.Println("购物完成", time.Since(start))
	//fmt.Println(moneyList)
}
