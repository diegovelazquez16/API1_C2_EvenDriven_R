package usecase

import (
	"holamundo/restaurantes/domain/models"
	"holamundo/restaurantes/domain/repository"
)

type UpdateRestauranteUseCase struct {
	RestauranteRepo repository.IRestauranteRepository
}

func (uc *UpdateRestauranteUseCase) Execute(restaurante *models.Restaurante) error {
	return uc.RestauranteRepo.Update(restaurante)
}