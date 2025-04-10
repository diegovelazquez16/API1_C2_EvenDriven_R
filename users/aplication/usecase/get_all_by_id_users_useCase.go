package usecase

import (
	"holamundo/users/domain/models"
	"holamundo/users/domain/repository"
)
type GetUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetUserUseCase) Execute(id uint) (*models.User, error) {
	return uc.UserRepo.GetByID(id)
}