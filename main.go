package main

import (
	"crud-carros-server2/infrastructure"
	"crud-carros-server2/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.ConnectDatabase()

	r := gin.Default()

	r.GET("/exhibir-carros", handlers.ExhibirCarros)
	r.DELETE("/comprar-carro/:id", handlers.ComprarCarro)

	r.Run(":8081")
}
