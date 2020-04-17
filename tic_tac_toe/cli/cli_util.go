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

func main() {
	fmt.Println("Tic Tac Toe , Welcome!\n What do you want to do? :)")
	playing := true
	for playing {
		selectedOpt := prompt.Input("> ", tabCompleter)
		if selectedOpt == "" {
			fmt.Println("nada por aqui nada por alla...")
		} else {
			fmt.Println("you selected " + selectedOpt)
		} //me devuelve un string "> selectedOpt"

		switch selectedOpt {
		case optNew:
			fmt.Println("Which size?") //int
			strNum := prompt.Input("new> ", newGameCompleter)
			num, _ := s.Atoi(strNum)
			myfun(num)
		case optPlay:
			fmt.Println("Play X or O")
			strings := prompt.Input("play> ", playCompleter)
			myfunc2(strings)
			fmt.Println("Choose coordinates. Example >> 0:1 Row first Column second")
			coor := prompt.Input("place> ", placeCompleter)
			//fmt.Printf(coor)
			if len(coor) == 0 {
				fmt.Printf("pero yo no veo nada!\n")
				continue
			}
			str := str.Split(coor, ":")
			myfunc3(str)
		case optExit:
			playing = false
		}
	}
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
		{Text: "O", Description: "player good"},
	}
	return prompt.FilterHasPrefix(s, player.GetWordBeforeCursor(), true)
}

func placeCompleter(player prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{{Text: "", Description: ""}}
	return prompt.FilterHasPrefix(s, player.GetWordAfterCursor(), true)
}

func myfun(num int) {
	fmt.Println("#,#,#\n,#,#,#\n,#,#,#")
	//call my func
}
func myfunc2(player string) {
	fmt.Printf("%s has played!\n", player)
	//call my func
	//prompt play until myfunc return winner
}
func myfunc3(coor []string) {

	fmt.Printf("mira que lindos estos dos numerettos %v %v\n", coor[0], coor[1])
}
