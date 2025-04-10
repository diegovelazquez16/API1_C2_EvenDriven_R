package usecase

import (
	"holamundo/restaurantes/domain/models"
	"holamundo/restaurantes/domain/repository"
)

type CreateRestauranteUseCase struct {
	RestauranteRepo repository.IRestauranteRepository
}

func (uc *CreateRestauranteUseCase) Execute(restaurante *models.Restaurante) error {
	return uc.RestauranteRepo.Create(restaurante)
}