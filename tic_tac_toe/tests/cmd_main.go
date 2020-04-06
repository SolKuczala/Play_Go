package main

import (
	"fmt"

	"../project/board"
)

var Board [3][3]string

func main() {
	Board = board.New()
	a := board.Coord{X: 0, Y: 0}
	b := board.Coord{X: 1, Y: 1}
	c := board.Coord{X: 2, Y: 2}
	err := board.Play("X", a, &Board)
	printError(err)
	err = board.Play("X", b, &Board)
	printError(err)
	err = board.Play("X", c, &Board)
	printError(err)
	board.PrintBoard(Board)
	printError(err)
	board.Check(&Board)
}

//funcion que haga la siguiente linea:
func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
