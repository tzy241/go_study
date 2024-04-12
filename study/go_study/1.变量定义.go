package main

import (
	"fmt"
	"goland/go_study/version"
)

var val = 13

func hello() {
	fmt.Println("hello world")
}

// 函数定义
func main() {
	//变量定义
	var name string
	name = "枫枫"
	fmt.Println(name)

	var name1 string = "枫枫"
	fmt.Println(name1)

	var name2 = "枫枫"
	fmt.Println(name2)

	name3 := "枫枫"
	fmt.Println(name3)

	hello()
	fmt.Println(name)

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	var (
		a3 = 3
		a4 = 4
	)
	fmt.Println(a3, a4)

	//常量定义
	const a5 = 5
	fmt.Println(a5)

	//跨文件引用
	fmt.Println(version.Version)
}
