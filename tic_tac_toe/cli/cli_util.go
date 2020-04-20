package main

import (
	"fmt"
	s "strconv"
	str "strings"

	"github.com/c-bata/go-prompt"
)

var optNew = "new"
var optPlay = "play"
var optExit = "exit"
var softExit = "exit main menu"
var hardExit = "exit game"

func getSizeFromUser() (int, error) {
	fmt.Println("Which size?") //int
	//pido el size del tablero
	strNum := prompt.Input("new> ", newGameCompleter)
	num, err := s.Atoi(strNum)
	if err != nil {
		fmt.Printf("We need numbers ('-.-)\n")
	}
	return num, err
}

func main() {
	fmt.Println("Tic Tac Toe , Welcome!\n What do you want to do? :)")
	//while para el juego entero
	playing := true

	for playing {
		selectedOpt := prompt.Input("begin> ", tabCompleter)
		if selectedOpt == "" {
			fmt.Println("nada por aqui nada por alla...")
		} else {
			fmt.Println("you selected " + selectedOpt + "\nPress tab to see options")
		} //me devuelve un string "> selectedOpt"

		switch selectedOpt {

		case optNew:
			needSize := true

			for needSize {
				size, err := getSizeFromUser()
				if err != nil {
					continue
				}
				myfun(size)
				needSize = false
			}

		case optPlay:
			//mientras no haya error o no haya ganador?
			missPlayer := true

			for missPlayer {
				fmt.Println("Play X or O")
				strings := prompt.Input("play> ", playCompleter)
				//checkear jugador
				if strings != "" {
					myfunc2(strings)
					missPlayer = false
				} else {
					fmt.Printf("At least a word man...\n")
					continue
				}
			}
			//una vez elegido el jugador>>
			missCoor := true

			var num1, num2 int
			var err error

			for missCoor {

				fmt.Println("Choose coordinates. Example >> 0:1 Row first, Column second")
				coor := prompt.Input("place> ", placeCompleter)
				//opcion de salir
				if coor == softExit {
					missCoor = false
				} else if coor == hardExit {
					missCoor = false
					playing = false
				}
				if len(coor) == 0 {
					fmt.Printf("pero yo no veo nada!\n")
					continue
				}
				//obtengo un array de dos cosas
				str := str.Split(coor, ":")
				//si tengo menos o mas de dos
				if len(str) < 2 || len(str) > 2 {
					fmt.Printf("Am I missing something? :/\n")
					continue
				}
				num1, err = s.Atoi(str[0])
				num2, err = s.Atoi(str[1])

				if err != nil {
					fmt.Printf("no es un numeretto >:[\n")
					continue
				} else {
					miraKlendosNumeretos(num1, num2)
					missCoor = false
				}
			}

		case optExit:
			playing = false
		} //end of switch

	} //end of game while
	fmt.Printf("Ci vediamo dopo")
}

func tabCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: optNew, Description: "New board Game"},
		{Text: optPlay, Description: "Play against other"},
		{Text: optExit, Description: "Exit the program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func newGameCompleter(numero prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "3", Description: "normal"},
		{Text: "4", Description: "well, complicated now"},
		{Text: "5", Description: "FO"},
	}
	return prompt.FilterHasPrefix(s, numero.GetWordBeforeCursor(), true)

}
func playCompleter(player prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "X", Description: "player Japan"},
		{Text: "O", Description: "player Sun"},
	}
	return prompt.FilterHasPrefix(s, player.GetWordBeforeCursor(), true)
}

func placeCompleter(player prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: softExit, Description: "To main menu"},
		{Text: hardExit, Description: "end program"}}
	return prompt.FilterHasPrefix(s, player.GetWordAfterCursor(), true)
}

func myfun(num int) {
	fmt.Println("#,#,#\n,#,#,#\n,#,#,#")
	//call my func
}
func myfunc2(player string) {
	fmt.Printf("%s selected!\n The other is the contrary :)\n", player)
	//call my func
	//prompt play until myfunc return winner
}
func miraKlendosNumeretos(num1, num2 int) bool {
	fmt.Printf("mira que lindos estos dos numerettos %v %v\n", num1, num2)
	fmt.Printf("con el turno puesto\n#,#,#\n,#,#,#\n,#,#,#\n")
	fmt.Printf("tu turno enemy\n")
	return false
	//call my func donde devuelve la jugadeta
}

func congrats() {
	fmt.Println("Congrats player!")
}
