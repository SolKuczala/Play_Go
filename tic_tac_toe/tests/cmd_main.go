package main

import (
	"fmt"

	"../project/board"
)

var Board [3][3]string

func main() {
	Board = board.New()
	plays := []string{
		"x", "o", "o",
		"o", "o", "x",
		"o", "x", "x",
	}
	for i, play := range plays {
		coor := board.Coord{X: uint(i % 3), Y: uint(i / 3)}
		err := board.Play(play, coor, &Board)
		printError(err)
		fmt.Printf("%s\n", board.Check(&Board))
		board.PrintBoard(Board)
	}
}

//funcion que haga la siguiente linea:
func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
