package main

import "fmt"

func main() {
	//乘法表
	// for i := 1; i <= 9; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// }

	//forrange
	// citys := [...]string{"北京", "上海", "成都"}
	// for i,v := range citys {
	// 	fmt.Printf("%d %s", i, v)
	// }

	//求数组[1, 3, 5, 7, 8]所有元素的和
	// arr := [...]int{1, 3, 5, 7, 8}
	// sum := 0
	// for _,v := range arr {
	// 	sum += v
	// }
	// fmt.Println(sum)

	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)

	arr := [3]int{1,2,3}
	fmt.Println(arr)
}
