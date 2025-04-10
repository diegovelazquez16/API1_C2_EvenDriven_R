package repository

import (
	"holamundo/pedidos/domain/models"
)

type IPedidoRepository interface {
	Create(pedido *models.Pedido) error
	GetAll() ([]models.Pedido, error)
	GetByID(id uint) (*models.Pedido, error)
	Update(pedido *models.Pedido) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}
