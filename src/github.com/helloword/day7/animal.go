package main

import "fmt"

//定义一个接口
type Animal interface {
	setName(string)
	getName() string
}

//定义一个结构体
type Dog struct {
	Name string
}

func (d *Dog) setName(name string) {
	d.Name = name
}

func (d Dog) getName() string {
	return d.Name
}

func main() {
	var wangcai = &Dog{
		Name: "wangcai",
	}

	var x Animal = wangcai
	fmt.Println(x.getName())

	x.setName("xiaohei")
	fmt.Println(x.getName())
}
