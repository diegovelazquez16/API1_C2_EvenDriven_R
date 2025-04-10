package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"holamundo/historial/aplication/usecase"
)

type HistorialDeleteController struct {
	DeleteHistorialUC *usecase.DeleteHistorialUseCase
}

func (c *HistorialDeleteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Historial ID"})
		return
	}

	err = c.DeleteHistorialUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Historial deleted successfully"})
}