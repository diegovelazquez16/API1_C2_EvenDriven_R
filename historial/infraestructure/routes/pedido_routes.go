package routes

import (
	
	"holamundo/historial/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func HistorialsRoutes(
	router *gin.Engine,
	createController *controllers.HistorialCreateController,
	getAllController *controllers.HistorialGetAllController,
	updateController *controllers.HistorialUpdateController,
	deleteController *controllers.HistorialDeleteController,
	getController *controllers.HistorialGetController,

) {
	group := router.Group("/Historials")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)  
		group.GET("/:id", getController.GetHistorialByID)   
		group.PUT("/:id", updateController.Update)    
		group.DELETE("/:id", deleteController.Delete )   
    
	}
}