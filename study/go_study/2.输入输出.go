package main

import "fmt"

func main() {
	fmt.Println("hello", "你好")
	fmt.Println("你好", "枫枫")
	fmt.Printf("%s 哇，你好美\n", "枫枫")
	fmt.Printf("%s 哇，你好美\n", "枫枫")
	fmt.Printf("%n\n", 3)
	fmt.Printf("%f\n", 3.1415)
	fmt.Printf("%T,%T\n", 3.1415, "你好")
	fmt.Printf("%v\n", "你好")
	fmt.Printf("%#v\n", "你好")

	var s = fmt.Sprintf("%v\n", "你好")
	fmt.Println(s)
}
