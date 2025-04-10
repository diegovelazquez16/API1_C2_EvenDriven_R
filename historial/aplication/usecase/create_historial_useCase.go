package usecase

import (
	"holamundo/historial/domain/models"
	"holamundo/historial/domain/repository"
)

type CreateHistorialUseCase struct {
	HistorialRepo repository.IHistorialRepository
}

func (uc *CreateHistorialUseCase) Execute(Historial *models.Historial) error {
	return uc.HistorialRepo.Create(Historial)
}
