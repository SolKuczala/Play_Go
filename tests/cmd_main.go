package main

import (
	"fmt"

	T "github.com/SolKuczala/tic-tac-go"
)

var gameStruct T.Game

func main() {
	gameStruct = T.NewGame(3)
	plays := []string{
		"o", "x", "x",
		"o", "x", "x",
		"o", "x", "x",
	}
	for i, play := range plays {
		coor := T.Coord{X: uint(i % 3), Y: uint(i / 3)}
		_, _, err := T.Play(play, coor, &gameStruct)
		printIfError(err)
		T.PrintBoard(&gameStruct)
	}
	T.PrintBoard(&gameStruct)
}

func printIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

//TODO: make the proper tests
