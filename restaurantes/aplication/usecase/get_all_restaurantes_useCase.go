package usecase

import (
	"holamundo/restaurantes/domain/models"
	"holamundo/restaurantes/domain/repository"
)
type GetAllRestaurantesUseCase struct {
	RestauranteRepo repository.IRestauranteRepository
}

func (uc *GetAllRestaurantesUseCase) Execute() ([]models.Restaurante, error) {
	return uc.RestauranteRepo.GetAll()
}

