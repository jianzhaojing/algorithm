package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

func main() {
	var user person
	user.name = "鼓励渣渣"
	user.city = "北京"
	user.age = 80

	fmt.Printf("渣渣的年纪为%v\n", user.age)
	fmt.Printf("渣渣的信息为%v\n", user)
	fmt.Printf("渣渣的信息为%#v\n", user)
}
