package tictactoe

import (
	"errors"
	"fmt"
	"strings"
)

const playerX string = "X"
const playerO string = "O"
const noPlayer string = "#"

var o, x = 0, 1

/*Game containing board type [][]string,	lastplayed type string,
winner type bool*/
type Game struct {
	Board      [][]string
	Lastplayed string
	Winner     bool
}

/*Coord Creates a struct for coordinates*/
type Coord struct {
	X uint
	Y uint
}

/*NewGame initializes and returns game struct*/
func NewGame(size int) Game {
	var matrix [][]string //[]

	for i := 0; i < size; i++ {
		matrix = append(matrix, []string{}) //append another slice, size times
		for j := 0; j < size; j++ {
			matrix[i] = append(matrix[i], noPlayer)
		}
	}
	var Game Game
	Game.Board = matrix
	Game.Lastplayed = noPlayer
	Game.Winner = false

	return Game
}

/*Play place the play, receiving the char, the coordinates
and the board itself. Returns string(char winner //or draw//), the board,
and error if:
game not available
place not available,
correct player,
correct turn,
TODO: draw*/
func Play(char string, coordinate Coord, game *Game) (string, [][]string, error) {
	if game.Winner == true {
		return "Last Game:", game.Board, errors.New("Cannot play, create another board to start again")
	}
	if game.Lastplayed == noPlayer || char != game.Lastplayed {
		err := checkAndPlace(char, coordinate, &game.Board)
		if err != nil {
			return "", game.Board, err
		}
		game.Lastplayed = char
	} else {
		return "", game.Board, errors.New("Hey, let the other play too :)")
	}

	return checkWinner(game), game.Board, nil
}

/*Called by Play.check if the player was correct and place the play otherwise returns errors:
* check if the place is available to play : else occupied error, no board error
 */
func checkAndPlace(char string, coordinate Coord, board *[][]string) error {
	x := int(coordinate.X) //column
	y := int(coordinate.Y) //row
	player := strings.ToUpper(char)
	//si la letra es la correspondiente
	if player != playerX && player != playerO {
		return errors.New("not a valid player")
	}
	matrix := *board
	//si las coordenadas no se pasan
	if l := len(*board); x < l && y < l {
		//si no esta ocupado, placea
		if matrix[y][x] == noPlayer {
			matrix[y][x] = player
		} else {
			return fmt.Errorf("Coordinate {%d %d} Occupied ! Try other place again", x, y)
		}
	} else {
		return fmt.Errorf("{%d %d} Oh oh, there is no board there my friend :/", x, y)
	}
	return nil
}

/*Check the game for winners. Used by Play, returns the player winner or empty(?*/
func checkWinner(game *Game) string {
	type conditions struct {
		row    [2]int
		column [2]int
		diag1  [2]int
		diag2  [2]int
	}
	var plays conditions
	o, x := 0, 1
	win := len(game.Board)

	for row := 0; row < len(game.Board); row++ {
		for column := 0; column < len(game.Board); column++ {

			//aca me voy a leer lo que viene por row y column(el string)
			matrix := game.Board
			rowPick := matrix[row][column]
			columnPick := matrix[column][row]
			if rowPick == noPlayer && columnPick == noPlayer {
				continue
				//esto va al column++
			}

			switch rowPick {
			case playerX:
				plays.row[x]++
			case playerO:
				plays.row[o]++
			}

			switch columnPick {
			case playerX:
				plays.column[x]++
			case playerO:
				plays.column[o]++
			}

			if row == column {
				if rowPick == playerX {
					plays.diag1[x]++
				} else {
					plays.diag1[o]++
				}
			}

			if row+column == len(game.Board)-1 {
				if rowPick == playerX {
					plays.diag2[x]++
				} else {
					plays.diag2[o]++
				}
			}

			//fmt.Printf("%+v %+v\n", row, column)

		} // end inner for

		//check si en row hay 3 iguales

		if plays.row[o] == win {
			game.Winner = true
			return playerO
		}
		plays.row[o] = 0

		if plays.row[x] == win {
			game.Winner = true
			return playerX
		}
		plays.row[x] = 0

		//check si column
		if plays.column[o] == win {
			game.Winner = true
			return playerO
		}
		plays.column[o] = 0

		if plays.column[x] == win {
			game.Winner = true
			return playerX
		}
		plays.column[x] = 0

		//check si diag1
		if plays.diag1[x] == win {
			game.Winner = true
			return playerX
		}
		if plays.diag1[o] == win {
			game.Winner = true
			return playerO
		}

		//check si diag2

		if plays.diag2[x] == win {
			game.Winner = true
			return playerX
		}
		if plays.diag2[o] == win {
			game.Winner = true
			return playerO
		}
	} //end of for
	return ""
}

/*PrintBoard prints a board nicely in cmd-line.*/
func PrintBoard(game *Game) {
	for i, innerArray := range game.Board {
		for j := range innerArray {
			fmt.Printf(game.Board[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("_")
}
