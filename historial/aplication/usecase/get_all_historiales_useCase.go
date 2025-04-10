package usecase

import (
	"holamundo/historial/domain/models"
	"holamundo/historial/domain/repository"
)
type GetAllHistorialsUseCase struct {
	HistorialRepo repository.IHistorialRepository
}

func (uc *GetAllHistorialsUseCase) Execute() ([]models.Historial, error) {
	return uc.HistorialRepo.GetAll()
}

