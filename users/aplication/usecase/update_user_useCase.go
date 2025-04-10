package usecase

import (
	"holamundo/users/domain/models"
	"holamundo/users/domain/repository"
)

type UpdateUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *UpdateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepo.Update(user)
}