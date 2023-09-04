package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxilovera/go-crud-example/dto"
	"github.com/maxilovera/go-crud-example/services"
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
	//obtener y validar parametros
	var personaActualizada dto.Persona

	if err := c.ShouldBindJSON(&personaActualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.ActualizarPersona(personaActualizada)

	if err != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Persona not found"})
		return
	}

	c.JSON(http.StatusTeapot, personaActualizada)
}

func EliminarPersona(c *gin.Context) {
	input := c.Param("id")

	id, err := strconv.Atoi(input)

	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Ingreso un id que no es numero entero"})
	}

	resultado := services.EliminarPersona(id)

	if (resultado != "") {
		c.JSON(http.StatusBadRequest, gin.H{"error" :"ID invalido"})
	} else {
		c.JSON(http.StatusAccepted, "correcto")
	}
}
