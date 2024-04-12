package main

import "fmt"

func main() {
	var arry = [3]int{1, 2, 3}
	fmt.Println(arry[0])
	arry[0] = 0
	fmt.Println(arry[0])

	var arry1 = []string{"a", "b", "c", "d"}
	fmt.Println(arry1)
}
