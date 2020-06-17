package main

import (
	"fmt"
	"strconv"
)

// "fmt"
// "strconv"

func main() {
	var i int = 20

	str := strconv.FormatInt(int64(i), 10)

	fmt.Printf("值：%s，类型：%T\n", str, str)
	fmt.Println(i)
}
