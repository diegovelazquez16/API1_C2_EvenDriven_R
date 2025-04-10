package usecase

import (
	"holamundo/pedidos/domain/models"
	"holamundo/pedidos/domain/repository"
)
type GetPedidosUseCase struct {
	PedidoRepo repository.IPedidoRepository
}

func (uc *GetPedidosUseCase) Execute(id uint) (*models.Pedido, error) {
	return uc.PedidoRepo.GetByID(id)
}