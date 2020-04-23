package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	T "../project/tictactoe"

	"github.com/c-bata/go-prompt"
)

const optNew = "new"
const optPlay = "play"
const softExit = "home"
const hardExit = "exit"

var gameStruct T.Game

//usar NewGame
//Play(juega y devuelve ganador)
//PrintBoard
func getSizeFromUser() (int, error) {
	fmt.Println("Which size?") //int
	//pido el size del tablero
	strNum := prompt.Input("new> ", exitCompleter)
	num, err := strconv.Atoi(strNum)
	if err != nil {
		err = errors.New("We need numbers ('-.-)\n")
	} else if num < 0 || num > 9 {
		err = errors.New("No negative numbers, and less than 10")
	}
	return num, err
}

func getPlayerFromUser() (string, error) {
	fmt.Println("Play X or O")
	strings := prompt.Input("play> ", playerCompleter)
	//checkear mas al jugador
	var err error
	if strings == "" {
		err = errors.New("at least a word man")
	}
	return strings, err
}

/*Execute get coordinates, returns slice of ints and error*/
func getCoorFromPlayer() (T.Coord, error) {
	var coorInt T.Coord
	fmt.Println("Choose coordinates with the following format >>\n0:1 Row first, Column second")
	//recibo string del player

	input := prompt.Input("place> ", exitCompleter)
	//si el string no existe
	if len(input) == 0 {
		return coorInt, errors.New("pero yo no veo nada!")
	}
	//obtengo un array de dos strings
	inputSplit := strings.Split(input, ":")
	//que no haya mas de un :, que no haya menos de dos strings
	if len(inputSplit) != 2 {
		return coorInt, errors.New("Am I missing something? :/\n")
	}
	//convierto los strings a num
	num1, err1 := strconv.Atoi(inputSplit[0])
	num2, err2 := strconv.Atoi(inputSplit[1])
	//si no son numeros, va error
	if err1 != nil || err2 != nil {
		return coorInt, errors.New("no es un numeretto >:[")
	}
	//si numero es neg, err
	if num1 < 0 || num2 < 0 {
		return coorInt, errors.New("no negative numbers, please")
	}
	coorInt = T.Coord{X: uint(num2), Y: uint(num1)}

	return coorInt, nil
	//TODO: ver si estoy checkeando que no pida numeros mayores a lo que el tablero me pide
}

func playAgain() bool {
	fmt.Println("Yes/No")
	decision := prompt.Input("play> ", exitCompleter)
	decision = strings.ToLower(decision)
	if decision == "yes" {
		return true
	}
	return false
}

func main() {
	fmt.Println("Tic Tac Toe , Welcome!\n What do you want to do? :)")
	//while para el juego entero
	playing := true

	for playing {
		selectedOpt := prompt.Input("begin> ", mainMenuCompleter)
		//menu
		switch selectedOpt {
		case optNew: //no tocar
			needSize := true

			for needSize {
				size, err := getSizeFromUser()
				if err != nil {
					fmt.Println(err)
					continue
				}
				myfunNewGame(size)
				needSize = false
			}
			//WORKING
		case optPlay:
			//lo voy a usar despues
			playerSelected := ""
			var oppositePlayer string
			var err error
			requirePlayer := true
			for requirePlayer {
				playerSelected, err = getPlayerFromUser()
				if err != nil {
					fmt.Println(err)
					continue
				}
				if playerSelected == "X" {
					oppositePlayer = "X"
				} else {
					oppositePlayer = "O"
				}
				requirePlayer = false
			}
			//una vez guardado los jugadores>>

			//mientras no haya error o no haya ganador
			var coor T.Coord
			var err2 error
			requireCoor := true
			for requireCoor {
				//check de cli
				coor, err2 = getCoorFromPlayer()
				if err != nil {
					fmt.Println(err2)
					continue
				}
				//guardada la coordenada >>
				player, winner, errgame := T.Play(playerSelected, coor, &gameStruct) //y board
				if winner {
					congrats(player)
					requireCoor = false

				} else {
					fmt.Println(errgame)
					continue
				}
			}

		case hardExit:
			playing = false

		case "":
			fmt.Println("nada por aqui nada por alla...")

		default:
			fmt.Println("you selected " + selectedOpt + "\nPress tab to see options")
			//me devuelve un string "> selectedOpt"

		} //end of switch

	} //end of while game
	fmt.Printf("Ci vediamo dopo")
}

func mainMenuCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: optNew, Description: "New board Game"},
		{Text: optPlay, Description: "Play against other"},
		{Text: hardExit, Description: "Exit the program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func playerCompleter(player prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "X", Description: "player Japan"},
		{Text: "O", Description: "player MG"},
	}
	return prompt.FilterHasPrefix(s, player.GetWordBeforeCursor(), true)
}

func exitCompleter(player prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: softExit, Description: "To main menu"},
		{Text: hardExit, Description: "Exit program"}}
	return prompt.FilterHasPrefix(s, player.GetWordAfterCursor(), true)
}

func myfunNewGame(num int) {
	fmt.Println("#,#,#\n,#,#,#\n,#,#,#")
	//call my func
}

func myFuncPlay(player string, coor [2]int) (string, bool, error) {
	fmt.Printf("mira que lindos estos dos numerettos %v %v\n", coor[0], coor[1])
	fmt.Printf("#,#,#\n,#,#,#\n,#,#,#\n\n")
	fmt.Printf("le toca al otro\n")
	return "", false, nil
	//call my func que recibe jugador, operacion y board, devuelve:char, game.board, nil/error
}

func congrats(player string) {
	fmt.Println("Congrats player %v!", player)
}
