package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int = 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//死循环
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
		break
	}
	//while模式
	var i int = 1
	for i <= 100 {
		sum += i
		i++
	}
	//遍历切片、map
	var s = []string{"枫枫", "知道"}
	for k, v := range s {
		fmt.Println(k, v)
	}
	var m = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for x := 0; x < 10; x++ {

	}
}
