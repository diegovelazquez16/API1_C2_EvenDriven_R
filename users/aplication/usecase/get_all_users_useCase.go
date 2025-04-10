package usecase

import (
	"holamundo/users/domain/models"
	"holamundo/users/domain/repository"
)
type GetAllUsersUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetAllUsersUseCase) Execute() ([]models.User, error) {
	return uc.UserRepo.GetAll()
}

