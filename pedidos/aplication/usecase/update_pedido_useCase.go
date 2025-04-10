package usecase

import (
	"holamundo/pedidos/domain/models"
	"holamundo/pedidos/domain/repository"
)

type UpdatePedidoUseCase struct {
	PedidoRepo repository.IPedidoRepository
}

func (uc *UpdatePedidoUseCase) Execute(user *models.Pedido) error {
	return uc.PedidoRepo.Update(user)
}