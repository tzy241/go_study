package main

import "fmt"

func main() {
	var age = 0
	fmt.Println("enter a:", age)
	fmt.Scan(&age)

	if age == 0 {
		fmt.Println("婴儿")
		return
	}
	if age <= 18 {
		fmt.Println("青年")
		return
	} else {
		return
	}
}
