package main

import "fmt"

type Person struct {
	name   string
	age    int
	sex    string
	height float32
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄：%v\n", p.name, p.age)
}

func (p *Person) setPerson(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	var p1 = Person{
		name: "lisi",
		age:  16,
		sex:  "男",
	}

	p1.PrintInfo()

	//修改结构体
	p1.setPerson("王五", 99)
	p1.PrintInfo()
}
