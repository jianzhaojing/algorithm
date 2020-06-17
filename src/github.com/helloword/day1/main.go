package main

import (
	"fmt"

	"strings"
)

func main() {
	a := 5;

	fmt.Printf("a=%d\n", a);
	// fmt.Println(unsafe.Sizeof(a));

	// fmt.Println("生命之源")

	b := "123-456-789";

	str := strings.Split(b, "-");
	fmt.Println(str);
}
