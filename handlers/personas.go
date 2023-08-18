package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-crud-example/dto"
	"github.com/go-crud-example/services"
)

func ObtenerPersonas(c *gin.Context) {
	//obtener y validar parametros
	//invocar a logica de negocio
	personas := services.ObtenerPersonas()

	//retornar la respuesta
	c.JSON(http.StatusOK, personas)
}

func ObtenerPersonaPorId(c *gin.Context) {
	//obtener y validar parametros
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro ID no numerico"})

		return
	}

	//invocar a logica de negocio
	persona := services.ObtenerPersonaPorId(id)

	//retornar la respuesta
	c.JSON(http.StatusOK, persona)
}

func CrearPersona(c *gin.Context) {
	//obtener y validar parametros
	var nuevaPersona dto.Persona
	if err := c.ShouldBindJSON(&nuevaPersona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//invocar a logica de negocio
	persona := services.CrearPersona(nuevaPersona)

	//retornar la respuesta
	c.JSON(http.StatusCreated, persona)
}

func ActualizarPersona(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro ID no numerico"})

		return
	}

	var nuevaPersona dto.Persona
	if err := c.ShouldBindJSON(&nuevaPersona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de persona invalidos"})
		return
	}

	persona := services.ActualizarPersona(id, &nuevaPersona)

	c.JSON((http.StatusOK), persona)
}

func EliminarPersona(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro ID no numerico"})
		return
	}

	persona := services.EliminarPersona(id)
	c.JSON((http.StatusOK), persona)

}
