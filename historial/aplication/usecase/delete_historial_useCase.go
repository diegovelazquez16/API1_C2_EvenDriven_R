package usecase

import (
	"holamundo/historial/domain/repository"
)
type DeleteHistorialUseCase struct {
	HistorialRepo repository.IHistorialRepository
}

func (uc *DeleteHistorialUseCase) Execute(id uint) error {
	return uc.HistorialRepo.Delete(id)
}