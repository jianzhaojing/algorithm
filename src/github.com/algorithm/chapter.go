package main

import "fmt"

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][2] = 2 //白子

	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d", v2)
		}
		fmt.Println()
	}

}
