package main

//validar la data que le llega y resp
import (
	"errors"
	"net/http"
	"strconv"

	T "github.com/SolKuczala/tic-tac-go"
	"github.com/gin-gonic/gin"
)

/*BOARD Main variable for saving the game*/
var BOARD T.Game

func main() {
	r := gin.Default()
	r.GET("/create-board/:size", createGame)
	//r.GET("/tossAcoin", tossCoin)
	r.PUT("/send-play/:player/:row/:column", sendPlay)
	r.GET("/status", getStatus)
	r.Run(":9090") // listen and serve on 0.0.0.0:9090
}

func createGame(c *gin.Context) {
	//tomo el size
	sizeParam := c.Param("size")
	//lo paso a numero
	size, err := strconv.Atoi(sizeParam)
	//si atoi no puede, devuelvo su error(horrible)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	//controlar size:
	if size < 0 || size > 9 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": ("No negative numbers, and less than 10")})
		return
	}
	BOARD = T.NewGame(size)
	c.JSON(http.StatusCreated, gin.H{"board": BOARD.Board})
	//c.JSON(200, gin.H{"message": "hola"})
	//T.PrintBoard(&board)
	//println(c) //prints this: 0xc00032e380
}

func sendPlay(c *gin.Context) {
	rowParam := c.Param("row")
	columnParam := c.Param("column")
	playerParam := c.Param("player")
	//checkeo player
	if playerParam != "X" && playerParam != "O" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": errors.New("X or O")})
		return
	}
	//una vez pasado eso, pasamos los numeros de str a int
	row, errR := strconv.Atoi(rowParam)
	column, errC := strconv.Atoi(columnParam)
	//checkeo errores
	if errR != nil || errC != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": errors.New("We need numbers ('-.-)\n")})
		return
	} else if row < 0 || row > 9 || column < 0 || column > 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errors.New("No negative numbers, and less than 10").Error()})
		return
	}
	//una vez que ta todo bien lo agregamos al struct con su corr format
	coor := T.Coord{X: uint(row), Y: uint(column)}
	//se lo pasamos a Play del package T
	winner, matrix, err := T.Play(playerParam, coor, &BOARD)
	//chekeamos si hay error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()}) //porq Error()
		return
	}
	//si hay winner
	if winner != "" {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "board": matrix, "winner": winner})
		//fmt.Println(BOARD)
		return
	}
	//sino muestro el estado para que se siga la partida
	c.JSON(http.StatusOK, gin.H{"status": "ok", "board": matrix, "byPlayer": BOARD.Lastplayed})
}

/*devuelve el board, a quien le toca*/
func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"board":       BOARD.Board,
		"last-player": BOARD.Lastplayed})
	return
}

/*see if i can make toss a coin option
func tossCoin(election bool)  {
	if election == tr{

	}
}*/
/*
proye:
que dos maquinas jueguen al tic tac toe.
una maquina va a crear el tablero
otra es la que pide? y empieza el juego
el p1 envia el codigo donde quiere que se marque su jugada
el p2 guarda eso y responde con la jugada siguiente
-----------
dos dockers que al levantarlos hacen la jugada
un externo cuenta los resultados, puede sacar estadisticas
*/
