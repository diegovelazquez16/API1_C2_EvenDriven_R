package usecase

import (
	"holamundo/pedidos/domain/repository"
)
type DeletePedidoUseCase struct {
	PedidoRepo repository.IPedidoRepository
}

func (uc *DeletePedidoUseCase) Execute(id uint) error {
	return uc.PedidoRepo.Delete(id)
}