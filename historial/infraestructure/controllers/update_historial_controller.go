package controllers

import (
	"net/http"
	"strconv"

	"holamundo/historial/aplication/usecase"
	"holamundo/historial/domain/models"

	"github.com/gin-gonic/gin"
)

type HistorialUpdateController struct {
	UpdateHistorialUC *usecase.UpdateHistorialUseCase
}

func (c *HistorialUpdateController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var historial models.Historial
	if err := ctx.ShouldBindJSON(&historial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	historial.ID = uint(id)

	if err := c.UpdateHistorialUC.Execute(&historial); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fallo al actualizar historial: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "historial actualizado correctamente",
		"historial":    historial,
	})
}