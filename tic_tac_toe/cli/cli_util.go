package main

import (
	"fmt"
	s "strconv"

	"github.com/c-bata/go-prompt"
)

var optNuevo = "nuevo"
var optJugar = "jugar"
var optSalir = "salir"

func main() {
	fmt.Println("Tic Tac Toe , Welcome!\n What do you want to do? :)")
	noSalir := true
	for noSalir {
		selectedOpt := prompt.Input("> ", tabCompleter)
		fmt.Println("You selected " + selectedOpt)
		//me devuelve un string "> selectedOpt"
		switch selectedOpt {
		case optNuevo:
			fmt.Println("Which size?") //int
			strNum := prompt.Input("> ", newgameCompleter)
			num, _ := s.Atoi(strNum)
			myfun(num)
		case optJugar:
			fmt.Println("Mandame una X o O")
			strings := prompt.Input("> ", playCompleter)
			myfunc2(strings)

		case optSalir:
			noSalir = false
		}
	}
	fmt.Printf("Ci vediamo dopo")

}

func newgameCompleter(numero prompt.Document) []prompt.Suggest {
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
		{Text: "O", Description: "player good"},
	}
	return prompt.FilterHasPrefix(s, player.GetWordBeforeCursor(), true)

}

func tabCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: optNuevo, Description: "New board Game"},
		{Text: optJugar, Description: "Play against other"},
		{Text: optSalir, Description: "Exit the program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func myfun(num int) {
	fmt.Println("mira que lindo %n", num)
}
func myfunc2(player string) {
	fmt.Println("%v has played! ", player)
	//call my func
	//prompt play until myfunc return winner
}
