package main

import "fmt"

func chvalue(n int) int {
	n = n + 1
	return n
}

func chvalue2(n *int) {
	*n = *n + 1
}

func main() {
	n := 10
	fmt.Println(n)
	chvalue(n)
	fmt.Println(n)
	chvalue2(&n)
	fmt.Println(n)
}
