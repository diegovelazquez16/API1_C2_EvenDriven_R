package usecase

import (
	"holamundo/restaurantes/domain/models"
	"holamundo/restaurantes/domain/repository"
)
type GetRestauranteUseCase struct {
	RestauranteRepo repository.IRestauranteRepository
}

func (uc *GetRestauranteUseCase) Execute(id uint) (*models.Restaurante, error) {
	return uc.RestauranteRepo.GetByID(id)
}	