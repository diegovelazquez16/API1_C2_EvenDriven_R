package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/pedidos/aplication/usecase"
)

type PedidoGetController struct {
	GetPedidoUC *usecase.GetPedidosUseCase
}

func (c *PedidoGetController) GetPedidoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe estar registrado o ser un numero valido"})
		return
	}

	pedido, err := c.GetPedidoUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, pedido)
}
