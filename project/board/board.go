package board

import (
	"errors"
	"fmt"
	"strings"
)

/*
*Initializes and returns an empty board.
 */
//TODO: cambiar el string a slice y el nombre
func NewBoard() [3][3]string {
	var matrix [3][3]string //[[, , ],[ , , ],[ , , ]]

	for i := 0; i < len(matrix); i++ {
		//me trae cada aarray de matrix
		for j := 0; j < len(matrix[i]); j++ {
			//me trae 1 valor de cada array
			matrix[i][j] = "#"
		}
	}
	//fmt.Printf("%+v", matrix)
	return matrix
}

/*
*Prints a board nicely in cmd-line.
 */
//TODO cambiar el nombre
func PrintBoard(board [3][3]string) {
	for i, innerArray := range board {
		for j, _ := range innerArray {
			fmt.Printf(board[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("_")
}

/*
*Makes the play receiving the letter, the coordinates for the placing
*and the board itself
 */
func Play(char string, coor Place, board *[3][3]string) error {
	//modifico el board segun la letra y la coordenada
	//este es el board vacio, check
	//coloco el string
	//devuelve board nuevo
	//faltaria auth
	x := int(coor.X)
	y := int(coor.Y)
	if l := len(board); x < l && y < l {
		if m := strings.ToUpper(char); m == "X" {
			if board[x][y] == "#" {
				board[x][y] = m
			} else {
				return errors.New(fmt.Sprintf("coordinate {%d %d} Occupied ! Try other coordinate again ", x, y))
			}
		} else {
			return errors.New("Hey, that' not an X my buddy")
		}
	} else {
		return errors.New(fmt.Sprintf("{%d %d} Oh oh, there is no board there my friend :/", x, y))
	}
	return nil
}

type Place struct {
	X uint
	Y uint
}

// other way to make types : type place [2]int
