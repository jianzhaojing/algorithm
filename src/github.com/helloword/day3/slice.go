package main

import "fmt" 

func main() {
	/*
	var a = []int{1,2,3,4,5};
	b := a[:];
	c := a[1:4];
	fmt.Println(b);
	fmt.Println(c);
	*/

	//make()创建切片
	var slice1 = make([]int, 4,8);
	fmt.Printf("slice1的值：%v，长度%d，容量%d", slice1, len(slice1), cap(slice1));
}