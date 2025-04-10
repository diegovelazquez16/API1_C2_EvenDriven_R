package routes

import (
	
	"holamundo/restaurantes/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func RestauranteRoutes(
	router *gin.Engine,
	createController *controllers.RestauranteCreateController,
	getAllController *controllers.RestauranteGetAllController,
	updateController *controllers.RestauranteUpdateController,
	deleteController *controllers.RestauranteDeleteController,
	getController *controllers.RestauranteGetController,

) {
	group := router.Group("/restaurante")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)  
		group.GET("/:id", getController.GetRestauranteByID)   
		group.PUT("/:id", updateController.Update)    
		group.DELETE("/:id", deleteController.Delete )   
    
	}
}