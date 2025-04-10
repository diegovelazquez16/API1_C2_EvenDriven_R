package usecase

import (
	"holamundo/restaurantes/domain/repository"
)
type DeleteRestauranteUseCase struct {
	RestauranteRepo repository.IRestauranteRepository
}

func (uc *DeleteRestauranteUseCase) Execute(id uint) error {
	return uc.RestauranteRepo.Delete(id)
}