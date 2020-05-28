package main

//state
import (
	"errors"
	"net/http"
	"strconv"

	T "github.com/SolKuczala/tic-tac-go"
	"github.com/gin-gonic/gin"
)

/*BOARD Main variable for saving the game*/
var GAME T.Game

func main() {
	r := gin.Default()
	r.GET("/create-board/:size", createGame)
	//r.GET("/tossAcoin", tossCoin)
	r.PUT("/send-play/:player/:row/:column", sendPlay)
	r.GET("/status", getStatus)
	r.Run(":9090") // listen and serve on 0.0.0.0:9090
}

func createGame(c *gin.Context) {
	sizeParam := c.Param("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	//controlar size:
	if size < 3 || size > 9 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": ("No negative numbers, and between 3 and 9")})
		return
	}
	GAME = T.NewGame(size)
	c.JSON(http.StatusCreated, gin.H{"board": GAME.Board})
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
	err := T.Play(playerParam, coor, &GAME)
	//chekeamos si hay error
	if err != nil {
		//error 400
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	//si hay winner
	if GAME.Status != T.GameStatusOngoing {
		if GAME.Status == T.GameStatusEndWithDraw {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "winner": "DRAW"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "winner": GAME.Lastplayed})
		}
		//fmt.Println(BOARD)
		return
	}
	//sino muestro el estado para que se siga la partida
	c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "byPlayer": GAME.Lastplayed})
}

/*devuelve el board, a quien le toca*/
func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"board":      GAME.Board,
		"lastPlayer": GAME.Lastplayed,
		"winners":    GAME.Status})
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
el primer jugador pide crear el tablero
el siguiente jugador pide el status y envia su jugada
el p2 pide status y responde con la jugada siguiente
-----------
dos dockers que al levantarlos hacen la jugada
TO-DO:un externo cuenta los resultados, puede sacar estadisticas
*/
