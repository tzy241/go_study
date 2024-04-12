package main

import "os"

func main() {
	file, err := os.OpenFile("go_study/文件/hello.txt", os.O_CREATE|os.WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("你好"))
}
