package main

import "fmt"

// 1.init()函数不能被其他函数调用， 而是再main函数执行之前自动被调用
// 2.init函数不能作为参数传入
// 3.不能有传入参数和返回值

func init() {
	fmt.Println("init-1")
}

func init() {
	fmt.Println("init-2")
}

func init() {
	fmt.Println("init-3")
}

func main() {

}
