package main

import "fmt"

func main() {
	type chapter struct {
		row int
		col int
		val int
	}

	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][2] = 2 //白子

	var newChessMap []chapter
	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2!=0 {
				valNode := chapter{
					row : i,
					col : j,
					val : v2,
				}

				newChessMap = append(newChessMap, valNode)
			}
		}
	}

	for _, chess := range newChessMap {
		fmt.Println(chess);
	}
}