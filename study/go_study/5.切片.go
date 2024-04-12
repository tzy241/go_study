package main

import (
	"fmt"
	"sort"
)

func main() {
	var list []string
	list = append(list, "枫枫")
	fmt.Println(list == nil)

	var list1 = []int{}
	list1 = make([]int, 3) //指定切片的长度
	fmt.Println(list1)

	//切片排序
	var lsit = []int{1, 2, 3, 4, 5, 8, 6}
	fmt.Println("排序前", lsit)
	sort.Ints(lsit)
	fmt.Println("升序", lsit)
	sort.Sort(sort.Reverse(sort.IntSlice(lsit)))
	fmt.Println("降序", lsit)

}
