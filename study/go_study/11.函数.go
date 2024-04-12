package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func checkId(id, username string) {
	fmt.Println(id, username)
}

func add(numberList ...int) {
	var sum int
	for _, num := range numberList {
		sum += num
	}
	fmt.Println("sum:", sum)
}

func rValue() int {
	return 5
}

func rValue2() (string, int) {
	return "", 1
}

func rValue3() (s string, ok bool) {
	if 1 < 2 {
		s = "李明"
		ok = true
	} else {
		s = "枫枫"
		ok = false
	}
	return
}

func main() {
	sayHello("李明")
	checkId("1", "1")
	add(1, 3, 4, 5, 6, 9)

	//不能在函数体中再创建函数，但是可以使用匿名函数,匿名函数和再外面定义没有区别
	var getName = func() string {
		return "张三"
	}
	fmt.Println(getName())

	//高阶函数
	fmt.Println("请输入你要执行的操作")
	fmt.Println("1 登录")
	fmt.Println("2 注册")
	fmt.Println("3 个人中心")
	var index int
	fmt.Scan(&index)

	//switch index {
	//case 1:
	//	login()
	//case 2:
	//	register()
	//case 3:
	//	userCenter()
	//}
	//等价于
	var funMap = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	//map[i]可以使用两个值进行接收,第一个值就是该值本身，第二个值是是否存在
	fun, ok := funMap[index]
	if ok {
		fun()
	}

	// 接收两个返回值
	var name, result = rValue3()
	fmt.Println(name, result)
	// 如果只想接收其中一个返回值，可以使用下划线来忽略不需要的值
	c, _ := rValue3()
	fmt.Println(c) // 输出：3
}
func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")
}
