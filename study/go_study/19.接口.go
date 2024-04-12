package main

import "fmt"

type Cat struct {
	Name string
}

func (this Cat) Sing() {
	fmt.Println(this.Name, "在唱歌")
}

type Chicken struct {
	Name string
}

func (this Chicken) Sing() {
	fmt.Println(this.Name, "在唱歌")
}

// 唱歌函数,但只支持Cat这一种结构体类型
func sing(c Cat) {
	c.Sing()
}

// 唱歌接口，接收所有实现了同一种方法的结构体
type SingInterface interface {
	Sing() //需要和结构体内部的方法都一致
}

// 接口类型变量可以接收 Cat类型变量也可以接收Dog类型变量
// 当结构体方法接收的是值类型，那么接口变量既可以接收值类型也可以接收指针类型
// 当结构体方法接收的是指针类型，那么接口变量只能接收指针类型
func animalSing(c SingInterface) {
	c.Sing()
	//类型断言，确定接口变量的具体类型
	ch, ok := c.(Chicken)
	fmt.Println(ch, ok)
	//第二种方式
	switch server := c.(type) {
	case Chicken:
		fmt.Println(server)
	case Cat:
		fmt.Println(server)
	default:
		fmt.Println("其他")
	}
}

// 空接口，可以接收任何结构体类型
type EmptyInterface interface {
}

func Print(val EmptyInterface) { //EmptyInterface可以使用any 或者 interface{}代替
	fmt.Println(val)
}

func main() {
	var ca = Cat{Name: "mimi"}
	sing(ca)
	var ch = Chicken{Name: "ik"}

	animalSing(ch)
	Print(ca)
	Print(ch)
}
