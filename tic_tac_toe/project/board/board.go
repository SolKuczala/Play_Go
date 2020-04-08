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
*Creates a struct for coordinates
**/
type Coord struct {
	X uint
	Y uint
}

/*
*Makes the play receiving the char, the coordinates for the placing
*and the board itself. Returns an error
 */

func Play(char string, coordinate Coord, board *[3][3]string) error {
	//modifico el board segun la letra y la coordenada
	//verifico que el segundo indice sea menor al board, coloco el string
	//devuelve board nuevo
	//faltaria auth
	x := int(coordinate.X) //column
	y := int(coordinate.Y) //row
	player := strings.ToUpper(char)
	//si la letra es la correspondiente
	if player != "X" && player != "O" {
		return errors.New(fmt.Sprintf("que quere vo"))
	}
	//si las coordenadas no se pasan
	if l := len(board); x < l && y < l {
		if board[y][x] == "#" {
			board[y][x] = player
		} else {
			return errors.New(fmt.Sprintf("coordinate {%d %d} Occupied ! Try other place again", x, y))
		}
	} else {
		return errors.New(fmt.Sprintf("{%d %d} Oh oh, there is no board there my friend :/", x, y))
	}
	return nil
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
			if row > 1 {
				if numberOf.column[x] == len(board) {
					return "Row! X won!"
				} else if numberOf.column[o] == len(board) {
					return "Row! O won!"
				} else {
					numberOf.column[x] = 0
					numberOf.column[o] = 0
				}
			}
			fmt.Printf("%+v %+v\n", row, column)
			rowCheck := board[row][column]
			columnCheck := board[column][row]
			if rowCheck == "#" && columnCheck == "#" {
				continue
				//esto sigue al siguiente column
			}

			if rowCheck == "X" || columnCheck == "X" {
				player = x
			} else if rowCheck == "O" || columnCheck == "O" {
				player = o
			}
			//el player que tenga la jugada...
			//11 10 01
			numberOf.column[player]++
			numberOf.row[player]++

			if row == column {
				numberOf.diag1[player]++
			}
			if row+column == len(board)-1 {
				numberOf.diag2[player]++
			}
			fmt.Printf("%+v\n", numberOf)
		}

		if numberOf.row[x] == len(board) {
			return "Row! X won!"
		} else if numberOf.row[o] == len(board) {
			return "Row! O won!"
		} else {
			numberOf.row[x] = 0
			numberOf.row[o] = 0
		}
	}

	if numberOf.diag1[x] == len(board) || numberOf.diag2[x] == len(board) {
		return "Diagonal! X won!"
	} else if numberOf.diag1[o] == len(board) || numberOf.diag2[o] == len(board) {
		return "Diagonal! O won!"
	}
	return "no win, play again?"
}

// other way to make types : type Coord [2]int
