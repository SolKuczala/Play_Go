package main

import (
	"fmt"

	T "github.com/SolKuczala/Play_Go/tic_tac_toe/tictactoe"

	"github.com/c-bata/go-prompt"
)

const optNew = "new"
const optPlay = "play"
const softExit = "home"
const hardExit = "exit"

var gameStruct T.Game

func main() {
	fmt.Println("Tic Tac Toe , Welcome!\n What do you want to do? :)")
	//while para el juego entero
	playing := true

	for playing {
		selectedOpt := prompt.Input("begin> ", mainMenuCompleter)
		//menu
		switch selectedOpt {
		case optNew:
			needSize := true

			for needSize {
				size, err := getSizeFromUser()
				if err != nil {
					fmt.Println(err)
					continue
				}
				gameStruct = T.NewGame(size)
				T.PrintBoard(&gameStruct)
				needSize = false
			}
			//WORKING
		case optPlay:
			//lo voy a usar despues
			playerSelected := ""
			var err error
			requirePlayer := true
			for requirePlayer {
				playerSelected, err = getPlayerFromUser()
				if err != nil {
					fmt.Println(err)
					continue
				}
				requirePlayer = false
			}
			//una vez guardado los jugadores>>

			//mientras no haya error o no haya ganador
			var coor T.Coord
			var err2 error
			var option string
			requireCoor := true
			for requireCoor {
				//check de cli
				coor, option, err2 = getCoorFromPlayer()
				if option == "home" {
					requireCoor = false
					break
				} else if option == "exit" {
					requireCoor = false
					playing = false
					break
				}
				if err2 != nil {
					fmt.Println(err2)
					continue
				}
				//guardada la coordenada >>
				//fmt.Printf("About to play: Player:%s, Coord:%+v\n", playerSelected, coor)
				winner, board, errgame := T.Play(playerSelected, coor, &gameStruct)
				//fmt.Printf("Result of play: Player:%s, ERR:%+v\n", winner, errgame)
				T.PrintBoard(&gameStruct)

				if errgame != nil {
					fmt.Println(errgame)
					continue
				}

				if winner == "" {
					gameStruct.Board = board
					playerSelected, err = switchPlayer(playerSelected)
					if err != nil {
						break
					} else {
						fmt.Printf("Your turn %s\n", playerSelected)
					}
				} else {
					congrats(winner)
					fmt.Println(board)
					gameStruct.Board = nil
					gameStruct.Lastplayed = ""
					requireCoor = false
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
	fmt.Printf("Ci vediamo dopo\n")
}

func congrats(player string) {
	fmt.Printf("Congrats player %v!", player)
}
