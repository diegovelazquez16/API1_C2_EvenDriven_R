package usecase

import (
	"holamundo/users/domain/repository"
)
type DeleteUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *DeleteUserUseCase) Execute(id uint) error {
	return uc.UserRepo.Delete(id)
}