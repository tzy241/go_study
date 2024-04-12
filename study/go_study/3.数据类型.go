package main

import (
	"fmt"
	"math"
)

func main() {
	var n1 uint8 = 255 //0-255
	fmt.Println(n1)
	// uint16
	// uint32
	// uint64
	// uint           默认等价与byte
	// int8           -128 - -1 、0 - 127
	//int16
	//int32
	//int64
	//int
	//int和uint的大小取决于所使用的平台
	var b byte = 255
	fmt.Printf("%T\n", b)

	var n2 float32 = math.MaxFloat32
	fmt.Println(n2)
	var n3 float64 = math.MaxFloat64
	fmt.Println(n3)
	//默认是float64

	var a byte = 'a'
	fmt.Printf("%c, %d\n", a, a)
	//多字节文字
	var z rune = '中'
	fmt.Printf("%c, %d\n", z, z)
	var s string = "你好"
	fmt.Printf("%s, %d\n", s, s)
	// ``原样输出字符串
	//bool值true和false 和其他类型无法相互转换
}
