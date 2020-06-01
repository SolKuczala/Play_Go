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

/*Game containing board type [][]string,lastplayed type string,
winner type bool*/
type Game struct {
	Board      [][]string
	Lastplayed string
	Status     int
}

var GameStatusOngoing = 0
var GameStatusEndWithWinner = 1
var GameStatusEndWithDraw = 2

/*Coord Creates a struct for coordinates*/
type Coord struct {
	X uint
	Y uint
}

/*NewGame initializes and returns game struct*/
func NewGame(size int) Game {
	var matrix [][]string

	for i := 0; i < size; i++ {
		matrix = append(matrix, []string{}) //append another slice, size times
		for j := 0; j < size; j++ {
			matrix[i] = append(matrix[i], noPlayer)
		}
	}
	var Game Game
	Game.Board = matrix
	Game.Lastplayed = noPlayer
	Game.Status = GameStatusOngoing

	return Game
}

/*Play :place the play, receiving the char, the coordinates
and the board itself. Returns string(char winner //or draw//), the board,
and error if:
game not available
place not available,
correct player,
correct turn,*/
func Play(char string, coordinate Coord, game *Game) error {
	if game.Status != GameStatusOngoing {
		return errors.New("Cannot play, create another board to start again")
	}
	if game.Lastplayed == noPlayer || char != game.Lastplayed {
		err := checkAndPlace(char, coordinate, &game.Board)
		if err != nil {
			return err
		}
		game.Lastplayed = char
	} else {
		return errors.New("Hey, let the other play too :)")
	}

	checkWinner(game)
	return nil
}

/*Called by Play.check if the player was correct and place the play otherwise returns errors:
* check if the place is available to play : else occupied error, no board error
 */
func checkAndPlace(char string, coordinate Coord, board *[][]string) error {
	x := int(coordinate.X) //row
	y := int(coordinate.Y) //column
	player := strings.ToUpper(char)
	//si la letra es la correspondiente
	if player != playerX && player != playerO {
		return errors.New("not a valid player")
	}
	matrix := *board
	//si las coordenadas no se pasan
	if l := len(*board); x < l && y < l {
		//si no esta ocupado, placea
		if matrix[x][y] == noPlayer {
			matrix[x][y] = player
		} else {
			return fmt.Errorf("Coordinate {%d %d} Occupied ! Try other place again", x, y)
		}
	} else {
		return fmt.Errorf("{%d %d} Oh oh, there is no board there my friend :/", x, y)
	}
	return nil
}

/*Check the game for winners. Used by Play, returns the player winner or empty(?*/
func checkWinner(game *Game) {
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

		} // end inner for
		//check si en row hay 3 iguales
		if plays.row[o] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.row[o] = 0

		if plays.row[x] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.row[x] = 0

		//check si column
		if plays.column[o] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.column[o] = 0

		if plays.column[x] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.column[x] = 0

		//check si diag1
		if plays.diag1[x] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.diag1[x] = 0

		if plays.diag1[o] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.diag1[o] = 0
		//check si diag2

		if plays.diag2[x] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.diag2[x] = 0
		if plays.diag2[o] == win {
			game.Status = GameStatusEndWithWinner
			return
		}
		plays.diag2[o] = 0
	} //end of for

	if game.Status == GameStatusOngoing {
		for _, row := range game.Board {
			for _, column := range row {
				if column == noPlayer {
					return
				}
			}
		}
		game.Status = GameStatusEndWithDraw
	}

	return
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
