package main

import "github.com/c-bata/go-prompt"

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
