package main

import "fmt"

// 函数awaitAdd返回的是一个匿名函数，该匿名函数接收切片参数
// go语言中函数可以返回另一个函数,返回后需要定义该函数
func awaitAdd(awaitSecond int) func(list ...int) (s int) {
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
