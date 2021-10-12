package main

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
	if size < 3 || size > 9 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": ("no negative numbers, and between 3 and 9")})
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
	if playerParam != "X" && playerParam != "O" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": errors.New("x or o")})
		return
	}
	row, errR := strconv.Atoi(rowParam)
	column, errC := strconv.Atoi(columnParam)
	if errR != nil || errC != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": errors.New("we need numbers ('-.-)\n")})
		return
	} else if row < 0 || row > 9 || column < 0 || column > 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errors.New("no negative numbers, and less than 10").Error()})
		return
	}
	coor := T.Coord{X: uint(row), Y: uint(column)}
	err := T.Play(playerParam, coor, &GAME)
	if err != nil {
		//error 400
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err.Error()})
		return
	}
	if GAME.Status != T.GameStatusOngoing {
		if GAME.Status == T.GameStatusEndWithDraw {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "winner": "DRAW"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "winner": GAME.Lastplayed})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "board": GAME.Board, "byPlayer": GAME.Lastplayed})
}

/*Returns the board status*/
func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"board":      GAME.Board,
		"lastPlayer": GAME.Lastplayed,
		"winners":    GAME.Status})
}

/* TO-DO: Some package that counts the results, to get stadistics */
