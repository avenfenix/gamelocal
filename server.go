// Parte que sirve los archivos de juegos. Es la "base de datos" en linea (local)
// La parte cliente le envia los archivos de juego actualizados
//

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("GameLocal version 0.0.1")
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
