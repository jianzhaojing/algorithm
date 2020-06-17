package main

import "fmt"

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() -", i)
	}
}
func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("mian() -", i)
	}
}
