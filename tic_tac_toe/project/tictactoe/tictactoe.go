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
and the board itself.returns an error(string) if incorrect player turn*/
func Play(char string, coordinate Coord, game *Game) (string, error) {
	//primer jugadeta si esta lleno de #: guardo el jugador
	//si ya tiene algo, compareta del char con lastplayed
	//
	if game.lastplayed == noPlayer || char != game.lastplayed {
		err := playCheck(char, coordinate, &game.board)
		if err != nil {
			return "", err
		}
		game.lastplayed = char
	} else {
		return "", errors.New("Hey, let the other play too :)")
	}
	result := check(game)
	if result != "" {
		fmt.Printf(result)
	}
	return "", nil
}

/*called by Play.check if the player was correct and place the play otherwise returns error*/
func playCheck(char string, coordinate Coord, board *[][]string) error {
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

/*Check the game for winners*/
func check(game *Game) string {
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
			return playerO
		}
		plays.row[o] = 0

		if plays.row[x] == win {
			return playerX
		}
		plays.row[x] = 0

		//check si column
		if plays.column[o] == win {
			return playerO
		}
		plays.column[o] = 0

		if plays.column[x] == win {
			return playerX
		}
		plays.column[x] = 0

		//check si diag1
		if plays.diag1[x] == win {
			return playerX
		}
		if plays.diag1[o] == win {
			return playerO
		}

		//check si diag2

		if plays.diag2[x] == win {
			return playerX
		}
		if plays.diag2[o] == win {
			return playerO
		}
	} //end of for
	return "no win, play again?"
}

/*fmt.Sprintf("%s wins", variable q contiene un string)*/
