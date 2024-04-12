package main

import "fmt"

// 函数awaitAdd返回的是一个匿名函数，该匿名函数接收切片参数
func awaitAdd(awaitSecond int) func(...int) (s int) {
	return func(numbers ...int) (sum int) {
		for _, i := range numbers {
			sum += i
		}
		return
	}
}

// 引用传参不就是传指针嘛
func set(name *string) {
	*name = "枫枫"
}

func main() {
	var fun1 = awaitAdd(3)
	var fun2 = fun1(1, 2, 3)
	fmt.Println(fun2)

	var s = "李华"
	fmt.Println(s)
	set(&s)
	fmt.Println(s)
}
