// 异常捕获
package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	var list = []int{1, 2}
	fmt.Println(list[2])
}

// 异常捕获函数
func readCatch() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	read()
}

func main() {
	readCatch()
}
