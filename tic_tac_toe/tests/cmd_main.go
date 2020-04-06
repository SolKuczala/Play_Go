package main

import (
	"fmt"

	"../project/board"
)

var Board [3][3]string

func main() {
	Board = board.New()
	c := board.Place{X: 0, Y: 2}
	err := board.Play("X", c, &Board)
	printError(err)
	board.PrintBoard(Board)
	printError(err)
}

//funcion que haga la siguiente linea:
func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
