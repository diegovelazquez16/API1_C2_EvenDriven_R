package launch

import (
	"holamundo/core"
	userUsecase "holamundo/users/aplication/usecase"
	userRepo "holamundo/users/domain/repository"
	userControllers "holamundo/users/infraestructure/controllers"
	userRoutes "holamundo/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserModule(router *gin.Engine) {
	userRepo := &userRepo.UserRepositoryImpl{DB: core.GetDB()}

	createUserUC := &userUsecase.CreateUserUseCase{UserRepo: userRepo}
	getAllUsersUC := &userUsecase.GetAllUsersUseCase{UserRepo: userRepo}
	getUserUC := &userUsecase.GetUserUseCase{UserRepo: userRepo}
	updateUserUC := &userUsecase.UpdateUserUseCase{UserRepo: userRepo}
	deleteUserUC := &userUsecase.DeleteUserUseCase{UserRepo: userRepo}
	loginUserUC := &userUsecase.LoginUserUseCase{UserRepo: userRepo}

	userCreateController := &userControllers.UserCreateController{CreateUserUC: createUserUC}
	userGetAllController := &userControllers.UserGetAllController{GetAllUsersUC: getAllUsersUC}
	userGetController := &userControllers.UserGetController{GetUserUC: getUserUC}
	userUpdateController := &userControllers.UserUpdateController{UpdateUserUC: updateUserUC}
	userDeleteController := &userControllers.UserDeleteController{DeleteUserUC: deleteUserUC}
	userLoginController := &userControllers.UserLoginController{LoginUseCase:loginUserUC}

	userRoutes.UserRoutes(router, userCreateController, userGetAllController, userUpdateController, userDeleteController, userGetController, userLoginController)
}
