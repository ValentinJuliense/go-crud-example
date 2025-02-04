package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud-example/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/personas", handlers.ObtenerPersonas)
	r.GET("/personas/:id", handlers.ObtenerPersonaPorId)
	r.POST("/personas", handlers.CrearPersona)
	r.PUT("/personas/:id", handlers.ActualizarPersona)
	r.DELETE("/personas/:id", handlers.EliminarPersona)

	r.Run(":8080")
}
