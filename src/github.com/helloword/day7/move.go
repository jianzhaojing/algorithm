package main

import "fmt"

//定义一个接口
type Animal interface {
	speck()
}

//定义一个结构体
type Dog struct {
}

func (d *Dog) speck() {
	fmt.Println("wangwangwang")
}
func main() {
	var animal Animal
	var dog = &Dog{}
	animal = dog
	animal.speck()
}
