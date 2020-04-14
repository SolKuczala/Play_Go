package board

import (
	"errors"
	"fmt"
	"strings"
)

var o, x = 0, 1

/*New initializes and returns an empty board.*/
//TODO: recibe dos numeros(o un struct)
func New(size int) [][]string {
	var matrix [][]string //[]

	for i := 0; i < size; i++ {
		matrix = append(matrix, []string{}) //append another slice, size times
		for j := 0; j < size; j++ {
			matrix[i] = append(matrix[i], "#")
		}
	}
	fmt.Println(matrix)

	return matrix
}

/*PrintBoard prints a board nicely in cmd-line.*/
//TODO cambiar el nombre
func PrintBoard(board [][]string) {
	for i, innerArray := range board {
		for j := range innerArray {
			fmt.Printf(board[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("_")
}

/*Coord Creates a struct for coordinates*/
type Coord struct {
	X uint
	Y uint
}

/*Play place the play, receiving the char, the coordinates
and the board itself. Returns an error if something goes wrong,otherwise nothing*/
func Play(char string, coordinate Coord, board *[][]string) error {
	//faltaria auth
	fmt.Println(board)
	fmt.Println(*board) //da el valor
	fmt.Println(&board) //da la dire

	x := int(coordinate.X) //column
	y := int(coordinate.Y) //row
	player := strings.ToUpper(char)
	//si la letra es la correspondiente
	if player != "X" && player != "O" {
		return errors.New("not the play")
	}
	matrix := *board
	fmt.Println(matrix)

	//si las coordenadas no se pasan
	if l := len(*board); x < l && y < l {
		if matrix[y][x] == "#" {
			matrix[y][x] = player
		} else {
			return fmt.Errorf("coordinate {%d %d} Occupied ! Try other place again", x, y)
		}
	} else {
		return fmt.Errorf("{%d %d} Oh oh, there is no board there my friend :/", x, y)
	}
	return nil
}

//Check podria llegar a tener un struct para el tipo de dato que recibe la matriz...para mejorar
//el segundo for
func Check(board *[][]string) string {
	//determina si alguien gana o no
	//
	type conditions struct {
		row    [2]int
		column [2]int
		diag1  [2]int
		diag2  [2]int
	}
	var plays conditions
	o, x := 0, 1
	win := len(*board)

	for row := 0; row < len(*board); row++ {
		for column := 0; column < len(*board); column++ {
			//primero voy a mirar 00 y despue27
			//row va ir por todos los rows
			//column va a ir por todas las columns

			//aca me voy a leer lo que viene por row y column(el string)
			matrix := *board
			rowPick := matrix[row][column]
			columnPick := matrix[column][row]
			if rowPick == "#" && columnPick == "#" {
				continue
				//esto va al column++
			}

			switch rowPick {
			case "X":
				plays.row[x]++
			case "O":
				plays.row[o]++
			}

			switch columnPick {
			case "X":
				plays.column[x]++
			case "O":
				plays.column[o]++
			}

			if row == column {
				if rowPick == "X" {
					plays.diag1[x]++
				} else {
					plays.diag1[o]++
				}
			}

			if row+column == len(*board)-1 {
				if rowPick == "X" {
					plays.diag2[x]++
				} else {
					plays.diag2[o]++
				}
			}

			fmt.Printf("%+v %+v\n", row, column)

		} // end inner for

		//check si en row hay 3 iguales

		if plays.row[o] == win {
			return "O wins"
		}
		plays.row[o] = 0

		if plays.row[x] == win {
			return "X wins"
		}
		plays.row[x] = 0

		//check si column
		if plays.column[o] == win {
			return "O wins"
		}
		plays.column[o] = 0

		if plays.column[x] == win {
			return "X wins"
		}
		plays.column[x] = 0

		//check si diag1
		if plays.diag1[x] == win {
			return "X wins"
		}
		if plays.diag1[o] == win {
			return "O wins"
		}

		//check si diag2

		if plays.diag2[x] == win {
			return "X wins"
		}
		if plays.diag2[o] == win {
			return "O wins"
		}
	} //end of for
	return "no win, play again?"
}
