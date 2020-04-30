package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	T "github.com/SolKuczala/Play_Go/tic_tac_toe/tictactoe"
	"github.com/c-bata/go-prompt"
)

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

/*Returns the selection and the error if exists*/
func getPlayerFromUser() (string, error) {
	fmt.Println("Play X or O")
	input := strings.ToUpper(prompt.Input("play> ", playerCompleter))
	var err error
	if input == "" {
		err = errors.New("at least a word man")
	}
	if input != "X" && input != "O" {
		err = errors.New("select player available at tab")
	}
	return input, err
}

/* Get coordinates, returns struct of coordinates, string if there was another option and error
*if incorrect input*/
func getCoorFromPlayer() (T.Coord, string, error) {
	var coorInt T.Coord
	fmt.Println("Choose coordinates with the following format >>\n0:1 Row first, Column second")

	//recibo string del player
	//outside options
	input := prompt.Input("place> ", exitCompleter)
	if input == "home" {
		return coorInt, softExit, nil
	}
	if input == "exit" {
		return coorInt, hardExit, nil
	}
	//si el string no existe
	if len(input) == 0 {
		return coorInt, "", errors.New("pero yo no veo nada!")
	}
	//obtengo un array de dos strings
	inputSplit := strings.Split(input, ":")
	//que no haya mas de un :, que no haya menos de dos strings
	if len(inputSplit) != 2 {
		return coorInt, "", errors.New("Am I missing something? :/\n")
	}
	//si uno de los 2 esta vacio
	if inputSplit[0] == "" || inputSplit[1] == "" {
		return coorInt, "", errors.New("one of them is an empty space ;/\n")
	}
	//convierto los strings a num
	num1, err1 := strconv.Atoi(inputSplit[0])
	num2, err2 := strconv.Atoi(inputSplit[1])
	//si no son numeros, va error
	if err1 != nil || err2 != nil {
		return coorInt, "", errors.New("no es un numeretto >:[")
	}
	//si numero es neg, err
	if num1 < 0 || num2 < 0 {
		return coorInt, "", errors.New("no negative numbers, please")
	}
	coorInt = T.Coord{X: uint(num2), Y: uint(num1)}

	return coorInt, "", nil
	//TODO: ver si estoy checkeando que no pida numeros mayores a lo que el tablero me pide
}

func switchPlayer(player string) (string, error) {
	if player == "X" {
		return "O", nil
	}
	if player == "O" {
		return "X", nil
	}
	return "", fmt.Errorf("Can't switch player %s", player)
}
