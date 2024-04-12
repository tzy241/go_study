package main

// 自定义数据类型
type Code int

// 自定义数据类型的意义,相当于给Code型变量提供一个 GetMsg的方法
func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	default:
		return ""
	}
}
func (c Code) OK() (code Code, msg string) {
	return c, c.GetMsg()
}

// 使用变量代替错误码
const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001
	NetworkErrCode Code = 1002
)

// 返回错误码和错误原因
func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.OK()
	}
	if name == "2" {
		return NetworkErrCode.OK()
	}
	if name == "3" {
		return SuccessCode.OK()
	}
}

func main() {

}
