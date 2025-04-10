package usecase

import (
	"holamundo/users/domain/models"
	"holamundo/users/domain/repository"
)

type CreateUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *CreateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepo.Create(user)
}