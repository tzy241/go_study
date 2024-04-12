// 向上抛
package main

import (
	"errors"
	"fmt"
)

// 底层函数
func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

// 逻辑函数
func server() (res int, err error) {
	res, err = div(2, 0)
	if err != nil {
		return
	}
	res++
	return
}

func main() {
	res, err := server()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
