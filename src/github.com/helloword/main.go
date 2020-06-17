package main

import "fmt"
var (
	name string
	age int
	isOk bool
)

func main() {
	if age := 19;age > 18 {
		fmt.Println("澳门赌场开业啦")
	} else {
		fmt.Println("快去写作业")
	}

	i :=0
	for ;i<10;i++ {
		fmt.Println(i)
	}
}