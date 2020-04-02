package main

import (
	"fmt"

	"../project/board"
)

var Board [3][3]string

func main() {
	Board = board.NewBoard()
	board.PrintBoard(Board)
	err := board.Play("X", 0, 1, &Board)
	printError(err)

	err = board.Play("X", 0, 1, &Board)
	printError(err)
	board.PrintBoard(Board)
}

//funcion que haga la siguiente linea:
func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
