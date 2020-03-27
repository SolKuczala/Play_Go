package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(400, gin.H{
			"message": "pong",
		})
		println(c) //prints this: 0xc00032e380
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
