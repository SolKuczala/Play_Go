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
func New() [3][3]string {
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
func Play(char string, coor Coord, board *[3][3]string) error {
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

type Coord struct {
	X uint
	Y uint
}

//podria llegar a tener un struct para el tipo de dato que recibe la matriz...para mejorar
//el segundo for
func Check(board *[3][3]string) string {
	//determina si alguien gana o no
	// TO DO: no me esta contando bien diag 2 y falta que cuente las rows
	type conditions struct {
		row    [2]int
		column [2]int
		diag1  [2]int
		diag2  [2]int
	}
	var player int
	var numberOf conditions
	o, x := 0, 1
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board); column++ {
			item := board[row][column]
			if item == "#" {
				continue
				//esto sigue al siguiente column
			}
			if item == "X" {
				player = x
			} else if item == "O" {
				player = o
			}

			numberOf.row[player]++

			if row == column {
				numberOf.diag1[player]++
			}
			if row+column == len(board)-1 {
				numberOf.diag2[player]++
			}
		}
	}
	fmt.Printf("%+v", numberOf)
	//if condiciones == 3 {
	//	return errors.Error("player tal win")
	//} else {
	//	return errors.Error("ninguno gano, va de nuevo?")
	//}
	return "dsaf"
}

// other way to make types : type Coord [2]int
