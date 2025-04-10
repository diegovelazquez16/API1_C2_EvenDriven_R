package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"holamundo/historial/aplication/usecase"
)
type HistorialGetAllController struct {
	GetAllHistorialUC *usecase.GetAllHistorialsUseCase
}
func (c *HistorialGetAllController) GetAll(ctx *gin.Context) {
	historial, err := c.GetAllHistorialUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, historial)
}