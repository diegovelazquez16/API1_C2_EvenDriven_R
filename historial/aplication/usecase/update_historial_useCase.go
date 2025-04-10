package usecase

import (
	"holamundo/historial/domain/models"
	"holamundo/historial/domain/repository"
)

type UpdateHistorialUseCase struct {
	HistorialRepo repository.IHistorialRepository
}

func (uc *UpdateHistorialUseCase) Execute(user *models.Historial) error {
	return uc.HistorialRepo.Update(user)
}