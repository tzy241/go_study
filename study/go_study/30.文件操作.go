package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	////一次性读取
	//byteData, err := os.ReadFile("go_study/文件/hello.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData))

	////分次读
	//file, err := os.Open("go_study/文件/hello.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close() //读取完毕关闭file
	//for {
	//	var byteData = make([]byte, 12)
	//	n, err := file.Read(byteData)
	//	//读取完毕
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(byteData), n)
	//}

	//按行读
	file, err := os.Open("go_study/文件/hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() //读取完毕关闭file

	//创建缓冲器
	buff := bufio.NewReader(file)
	for {
		line, _, err := buff.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	//按分隔符

}
