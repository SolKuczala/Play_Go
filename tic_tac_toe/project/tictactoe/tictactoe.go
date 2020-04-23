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

/*Game containin board type [][]string,	lastplayed type string,
winner type bool*/
type Game struct {
	board      [][]string
	lastplayed string
	winner     bool
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
	fmt.Println(matrix)
	var Game Game
	Game.board = matrix
	Game.lastplayed = noPlayer
	Game.winner = false

	return Game
}

/*PrintBoard prints a board nicely in cmd-line.*/
func PrintBoard(game *Game) {
	for i, innerArray := range game.board {
		for j := range innerArray {
			fmt.Printf(game.board[i][j])
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
and the board itself. Returns string(char winner or draw), the board, and error if:
*place not available, correct player, correct turn, draw
*/
func Play(char string, coordinate Coord, game *Game) (string, [][]string, error) {
	if game.lastplayed == noPlayer || char != game.lastplayed {
		err := placeCheck(char, coordinate, &game.board)
		if err != nil {
			return "", game.board, err
		}
		game.lastplayed = char
	} else {
		return char, game.board, errors.New("Hey, let the other play too :)")
	}

	gameState, err := check(game)
	if err != nil {
		return "", game.board, err
	}
	return gameState, game.board, nil
}

/*called by Play.check if the player was correct and place the play otherwise returns errors:
* check if the place is available to play : else occupied error, no board error
 */
func placeCheck(char string, coordinate Coord, board *[][]string) error {
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
		if matrix[y][x] == noPlayer {
			matrix[y][x] = player
		} else {
			return fmt.Errorf("coordinate {%d %d} Occupied ! Try other place again", x, y)
		}
	} else {
		return fmt.Errorf("{%d %d} Oh oh, there is no board there my friend :/", x, y)
	}
	return nil
}

/*Check the game for winners. Used by Play, returns the player winner or an error if draw*/
func check(game *Game) (string, error) {
	type conditions struct {
		row    [2]int
		column [2]int
		diag1  [2]int
		diag2  [2]int
	}
	var plays conditions
	o, x := 0, 1
	win := len(game.board)

	for row := 0; row < len(game.board); row++ {
		for column := 0; column < len(game.board); column++ {

			//aca me voy a leer lo que viene por row y column(el string)
			matrix := game.board
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

			if row+column == len(game.board)-1 {
				if rowPick == playerX {
					plays.diag2[x]++
				} else {
					plays.diag2[o]++
				}
			}

			fmt.Printf("%+v %+v\n", row, column)

		} // end inner for

		//check si en row hay 3 iguales

		if plays.row[o] == win {
			return playerO, nil
		}
		plays.row[o] = 0

		if plays.row[x] == win {
			return playerX, nil
		}
		plays.row[x] = 0

		//check si column
		if plays.column[o] == win {
			return playerO, nil
		}
		plays.column[o] = 0

		if plays.column[x] == win {
			return playerX, nil
		}
		plays.column[x] = 0

		//check si diag1
		if plays.diag1[x] == win {
			return playerX, nil
		}
		if plays.diag1[o] == win {
			return playerO, nil
		}

		//check si diag2

		if plays.diag2[x] == win {
			return playerX, nil
		}
		if plays.diag2[o] == win {
			return playerO, nil
		}
	} //end of for
	return "", errors.New("no win, play again?")
}

/*fmt.Sprintf("%s wins", variable q contiene un string)*/
