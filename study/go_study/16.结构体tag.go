package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` //忽略为空值,age值为0时不显示
	Password string `json:"-"`             //- 不显示
}

func main() {
	var user = User{Name: "枫枫", Age: 0, Password: "123456"}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}
