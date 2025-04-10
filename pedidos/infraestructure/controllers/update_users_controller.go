package controllers

import (
	"net/http"
	"strconv"

	"holamundo/pedidos/aplication/usecase"
	"holamundo/pedidos/domain/models"

	"github.com/gin-gonic/gin"
)

type PedidoUpdateController struct {
	UpdatePedidoUC *usecase.UpdatePedidoUseCase
}

func (c *PedidoUpdateController) Update(ctx *gin.Context) {
	// Obtener el ID de la URL y validarlo
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// Mapear el JSON del cuerpo de la solicitud a la estructura User
	var pedido models.Pedido
	if err := ctx.ShouldBindJSON(&pedido); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	// Asignar el ID de la URL al objeto User
	pedido.ID = uint(id)

	// Ejecutar el caso de uso para actualizar el usuario
	if err := c.UpdatePedidoUC.Execute(&pedido); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fallo al actualizar pedido: " + err.Error(),
		})
		return
	}

	// Devolver una respuesta exitosa con el usuario actualizado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Pedido updated successfully",
		"pedido":    pedido,
	})
}