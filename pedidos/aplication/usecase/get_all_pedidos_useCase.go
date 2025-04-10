package usecase

import (
	"holamundo/pedidos/domain/models"
	"holamundo/pedidos/domain/repository"
)
type GetAllPedidosUseCase struct {
	PedidoRepo repository.IPedidoRepository
}

func (uc *GetAllPedidosUseCase) Execute() ([]models.Pedido, error) {
	return uc.PedidoRepo.GetAll()
}

