package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("婴儿")
	case age <= 18:
		fmt.Println("青少年")
		fallthrough //继续往下走
	case age <= 25:
		fmt.Println("成年")
	default:
		fmt.Println("老年")
	}

}
