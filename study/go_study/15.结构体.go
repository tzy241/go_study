package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
}

// 给结构体添加方法
func (s Student) Study() {
	fmt.Printf("%s正在学习！\n", s.Name)
}

// 组合
func (s Student) Info() {
	fmt.Printf("名字：%s 班级：%s\n", s.Name, s.Class.Name)
}

// 修改结构体内部的值
func (s *Student) SetName(name string) {
	s.Name = name
}

type Test struct {
	Age  int
	Name string
}

func main() {
	var s1 = Student{Name: "枫枫"}
	s1.Study()

	var c1 = Class{Name: "三年一班"}
	var s2 = Student{Class: c1, Name: "枫枫"}
	s2.Info()

	s2.SetName("李华")
	s2.Info()

	var y = Test{12, "li"}
	fmt.Println(y)
}
