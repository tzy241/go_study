package main

import "fmt"

//1.这些调用直到return前才被执行，因此，可以用来做资源清理
//2.多个defer语句，按照先进后出的方式执行，谁离return近谁先执行
//3.defer语句中的变量，再defer声明时就决定了

func main() {
	defer fmt.Println("defer-1")
	defer fmt.Println("defer-2")
	defer func() {
		fmt.Println("defer-3")
		//函数内部只能拿到已经声明的变量
	}()

	var name = "枫枫"
	defer fmt.Println(name)
}
