package usecase

import (
	"holamundo/historial/domain/models"
	"holamundo/historial/domain/repository"
)
type GetHistorialsUseCase struct {
	HistorialRepo repository.IHistorialRepository
}

func (uc *GetHistorialsUseCase) Execute(id uint) (*models.Historial, error) {
	return uc.HistorialRepo.GetByID(id)
}