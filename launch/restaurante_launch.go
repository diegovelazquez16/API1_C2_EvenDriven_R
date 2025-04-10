package launch

import (
	"holamundo/core"
	restauranteUsecase "holamundo/restaurantes/aplication/usecase"
	restauranteRepo "holamundo/restaurantes/domain/repository"
	restauranteControllers "holamundo/restaurantes/infraestructure/controllers"
	restauranteRoutes "holamundo/restaurantes/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRestauranteModule(router *gin.Engine) {
	restauranteRepo := &restauranteRepo.RestauranteRepositoryImpl{DB: core.GetDB()}

	createRestauranteUC := &restauranteUsecase.CreateRestauranteUseCase{RestauranteRepo: restauranteRepo}
	getAllRestaurantesUC := &restauranteUsecase.GetAllRestaurantesUseCase{RestauranteRepo: restauranteRepo}
	getRestauranteUC := &restauranteUsecase.GetRestauranteUseCase{RestauranteRepo: restauranteRepo}
	updateRestauranteUC := &restauranteUsecase.UpdateRestauranteUseCase{RestauranteRepo: restauranteRepo}
	deleteRestauranteUC := &restauranteUsecase.DeleteRestauranteUseCase{RestauranteRepo: restauranteRepo}

	restauranteCreateController := &restauranteControllers.RestauranteCreateController{CreateRestauranteUC: createRestauranteUC}
	restauranteGetAllController := &restauranteControllers.RestauranteGetAllController{GetAllRestaurantesUC: getAllRestaurantesUC}
	restauranteGetController := &restauranteControllers.RestauranteGetController{GetRestauranteUC: getRestauranteUC}
	restauranteUpdateController := &restauranteControllers.RestauranteUpdateController{UpdateRestauranteUC: updateRestauranteUC}
	restauranteDeleteController := &restauranteControllers.RestauranteDeleteController{DeleteRestauranteUC: deleteRestauranteUC}

	restauranteRoutes.RestauranteRoutes(router, restauranteCreateController, restauranteGetAllController, restauranteUpdateController, restauranteDeleteController, restauranteGetController)
}
