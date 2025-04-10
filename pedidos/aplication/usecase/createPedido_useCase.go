package usecase

import (
	"holamundo/pedidos/domain/models"
	"holamundo/pedidos/domain/repository"
)

type CreatePedidoUseCase struct {
	PedidoRepo repository.IPedidoRepository
}

func (uc *CreatePedidoUseCase) Execute(pedidos *models.Pedido) error {
	return uc.PedidoRepo.Create(pedidos)
}
