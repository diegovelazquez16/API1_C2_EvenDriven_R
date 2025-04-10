package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"holamundo/historial/aplication/usecase"
	"holamundo/historial/domain/models"
)

type HistorialCreateController struct {
	CreateHistorialUC *usecase.CreateHistorialUseCase
}	

func (c *HistorialCreateController) Create(ctx *gin.Context) {
	var historial models.Historial
	if err := ctx.ShouldBindJSON(&historial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreateHistorialUC.Execute(&historial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Historial creado de forma exitosa",
		"Historial": historial,
	})
}
