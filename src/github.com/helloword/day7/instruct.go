package main

import "fmt"

//定义一个接口
type Computer interface {
	start()
	stop()
}

//定义一个手机结构体
type Phone struct {
}

//手机的方法
func (p Phone) start() {
	fmt.Println("手机开机")
}

func (p Phone) stop() {
	fmt.Println("手机关机")
}

//定义一个相机结构体
type Camera struct {
}

//相机的方法
func (c Camera) start() {
	fmt.Println("相机开机")
}

func (c Camera) stop() {
	fmt.Println("相机关机")
}

func main() {
	var phone = Phone{}

	var computer Computer = phone
	computer.start()
	computer.stop()

	var camera = Camera{}
	computer = camera
	computer.start()
	computer.stop()
}
