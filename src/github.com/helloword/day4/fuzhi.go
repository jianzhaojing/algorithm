package main 

import "fmt"

func main() {
	x := 1;
	y := x;

	fmt.Printf("x的地址为：%v\n,y的地址为:%v\n", &x, &y);
}