//中断程序，一般用于初始化

package main

import (
	"fmt"
	"os"
)

func init() {
	_, err := os.ReadFile("1234")
	if err != nil {
		panic("出错了")
	}
}

func main() {
	fmt.Println("Hello World")
}
