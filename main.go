package main

import (
	"fmt"
)

var (
	board  *[][]int
	witdh  = 10
	height = 10
)

func main() {
	initializeBoard()
	drawBoard()
}

func initializeBoard() {
	for i := 0; i < height; i++ {

		if board == nil {
			temp := make([][]int, 0)
			board = &temp
		}

		row := make([]int, witdh)
		*board = append(*board, row)
	}
}

func drawBoard() {
	strColumnIndexes := "  "
	for column, _ := range (*board)[0] {
		strColumnIndexes = fmt.Sprintf("%s %d", strColumnIndexes, column)
	}

	fmt.Println(strColumnIndexes)

	for numRow, row := range *board {
		fmt.Println(numRow, row)
	}
}
