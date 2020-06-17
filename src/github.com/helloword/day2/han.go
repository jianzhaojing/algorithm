package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "hello沙河小王子"
	count := 0
	for _, v := range str {
		// fmt.Printf("%c\n", v)
		// if unicode.Is(unicode.Scripts["Han"], v) {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}

	fmt.Println(count)
}
