package routes

import (
	
	"holamundo/pedidos/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func PedidosRoutes(
	router *gin.Engine,
	createController *controllers.PedidoCreateController,
	getAllController *controllers.PedidoGetAllController,
	updateController *controllers.PedidoUpdateController,
	deleteController *controllers.PedidoDeleteController,
	getController *controllers.PedidoGetController,

) {
	group := router.Group("/pedidos")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)  
		group.GET("/:id", getController.GetPedidoByID)   
		group.PUT("/:id", updateController.Update)    
		group.DELETE("/:id", deleteController.Delete )   
    
	}
}