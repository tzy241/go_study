package main

// T可以接收uint 和 int 型 ，同类型
func plus[T uint | int](n1, n2 T) T {
	return n1 + n2
}

// 多类型
func myPrint[T int, K string | int](n1 T, n2 K) {

}

// 提炼多类型
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func plus1[T Number](n1, n2 T) T {
	return n1 + n2
}

// 泛型结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func main() {
	var u1, u2 = uint(2), uint(3)
	println(plus(u1, u2))

}
