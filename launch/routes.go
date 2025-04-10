package launch

import (
"github.com/gin-gonic/gin"
"holamundo/pedidos/infraestructure/messaging"


)



func RegisterRoutes(router *gin.Engine, pedidoPublisher * messaging.PedidoPublisher) {
	RegisterUserModule(router)
	RegisterRestauranteModule(router)
	RegisterPedidoModule(router, pedidoPublisher)
	RegisterHistorialModule(router)
}
