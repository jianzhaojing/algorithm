package main

import "fmt"

func main() {
	// var slice1 = []string{"php", "java"};
	// var slice2 = []string{"c", "c++"};
	// var slice3 []string

	// slice3 = append(slice1,slice2...);

	// fmt.Println(slice3);

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	
}