package main

import "fmt"

//1.自定义数据类型可以绑定方法，而类型别名不能
//2.打印出的还是原始类型
//3.和原始类型比较，类型别名不用转换，而自定义数据类型需要转换

type MyCode int
type myCode = int

const a MyCode = 1
const b myCode = 1

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	var age = 1
	if b == age {
		//类型别名可以直接与age进行判断
	}
	if int(a) == age {
		//自定义数据类型判断需要进行转换
	}
}
