package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"holamundo/pedidos/aplication/usecase"
)
type PedidoGetAllController struct {
	GetAllPedidosUC *usecase.GetAllPedidosUseCase
}
//cambia GetAll a execut o run y lo mismo para todos
func (c *PedidoGetAllController) GetAll(ctx *gin.Context) {
	pedidos, err := c.GetAllPedidosUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pedidos)
}