package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/create-game", createGame)
	r.POST("/send-play", sendPlay)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}

func createGame(c *gin.Context) {
	//crea el board de juego que le manda al user cuando va al path /create-game
	c.JSON(200, gin.H{"message": "hola"})

	println(c) //prints this: 0xc00032e380
}

func sendPlay(c *gin.Context) {
	//tengo que recibir
	b := make([]byte, 2048)
	r, _ := c.Request.Body.Read(b)
	fmt.Printf("Me llego:\n%+v", r)
}

/*proye:
que dos maquinas jueguen al tic tac toe.
una maquina va a crear el tablero
otra es la que pide y empieza el juego
el p1 envia el codigo donde quiere que se marque su jugada
el p2 guarda eso y responde con la jugada siguiente
-----------
dos dockers que al levantarlos hacen la jugada
un externo cuenta los resultados, puede sacar estadisticas
*/
