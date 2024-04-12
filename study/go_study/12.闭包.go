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

// 闭包示例1.,闭包函数中注意变量名不要重复使用，函数名中如果没有指定变量名此时返回值不需要加括号，返回时就需要指定要返回的变量
func addNumbers(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 闭包示例2：返回两个函数的函数
func testTwoReturn(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
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

	tmp1 := addNumbers(100)
	fmt.Println(tmp1(1), tmp1(2), tmp1(3))
	f1, f2 := testTwoReturn(10)
	fmt.Println(f1(1), f2(1))
}
