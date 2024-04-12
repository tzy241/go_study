package main

import "fmt"

func main() {
	var userMap = map[int]string{
		1: "枫枫",
		2: "小明",
		4: "",
	}
	fmt.Println(userMap)
	fmt.Println(userMap[2])
	fmt.Printf("%#v \n", userMap[2])

	var value = userMap[4]
	fmt.Println(value)

	//使用两个值去接受map中的某个元素，第二值表示该值是否存在，第一值接收
	var value2, ok = userMap[2]
	fmt.Println(value2, ok)

	delete(userMap, 4)
	fmt.Println(userMap)
	//map声明时需要初始化
	var map1 = make(map[int]string)
	fmt.Println(map1)
	var map2 = map[int]string{}
	fmt.Println(map2)
}
