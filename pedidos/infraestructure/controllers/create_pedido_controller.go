package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"holamundo/pedidos/aplication/usecase"
	"holamundo/pedidos/domain/models"
	"holamundo/pedidos/infraestructure/messaging"
)

type PedidoCreateController struct {
	CreatePedidoUC  *usecase.CreatePedidoUseCase
	PedidoPublisher *messaging.PedidoPublisher
}

func (c *PedidoCreateController) Create(ctx *gin.Context) {
	var pedido models.Pedido
	if err := ctx.ShouldBindJSON(&pedido); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreatePedidoUC.Execute(&pedido)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el publicador está inicializado
	if c.PedidoPublisher == nil {
		log.Println("⚠️ Pedido creado, pero PedidoPublisher no está inicializado")
		ctx.JSON(http.StatusAccepted, gin.H{"warning": "Pedido creado, pero no se pudo enviar a RabbitMQ porque PedidoPublisher no está inicializado"})
		return
	}

	// Convertir a PedidoQueue y publicarlo en RabbitMQ
	pedidoQueue := messaging.PedidoQueue{
		ID:            pedido.ID,
		UsuarioID:     pedido.UsuarioID,
		RestauranteID: pedido.RestauranteID,
		Producto1:     pedido.Producto1,
		Producto2:     pedido.Producto2,
		Producto3:     pedido.Producto3,
		Total:         pedido.Total,
		Estado:        pedido.Estado,
		PagoID:        pedido.PagoID,
	}

	err = c.PedidoPublisher.Publish(pedidoQueue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el mensaje a la cola"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Pedido creado y enviado a la cola de RabbitMQ",
		"Pedido":  pedido,
	})
}
