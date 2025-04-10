package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"holamundo/pedidos/aplication/usecase"
)

type PedidoDeleteController struct {
	DeletePedidoUC *usecase.DeletePedidoUseCase
}

func (c *PedidoDeleteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pedido ID"})
		return
	}

	err = c.DeletePedidoUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pedido deleted successfully"})
}