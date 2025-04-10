package launch

import (
	"holamundo/core"
	historialUsecase "holamundo/historial/aplication/usecase"
	historialRepo "holamundo/historial/domain/repository"
	historialControllers "holamundo/historial/infraestructure/controllers"
	historialRoutes "holamundo/historial/infraestructure/routes"
	messaging "holamundo/historial/infraestructure/messaging"
	"log"
	"github.com/gin-gonic/gin"
)
//historial
func RegisterHistorialModule(router *gin.Engine) {
	historialRepo := &historialRepo.HistorialRepositoryImpl{DB: core.GetDB()}

	createHistorialUC := &historialUsecase.CreateHistorialUseCase{HistorialRepo: historialRepo}
	getAllhistorialUC := &historialUsecase.GetAllHistorialsUseCase{HistorialRepo: historialRepo}
	getHistorialUC := &historialUsecase.GetHistorialsUseCase{HistorialRepo: historialRepo}
	updateHistorialUC := &historialUsecase.UpdateHistorialUseCase{HistorialRepo: historialRepo}
	deleteHistorialUC := &historialUsecase.DeleteHistorialUseCase{HistorialRepo: historialRepo}

	historialCreateController := &historialControllers.HistorialCreateController{CreateHistorialUC: createHistorialUC}
	historialGetAllController := &historialControllers.HistorialGetAllController{GetAllHistorialUC: getAllhistorialUC}
	historialGetController := &historialControllers.HistorialGetController{GetHistorialUC: getHistorialUC}
	historialUpdateController := &historialControllers.HistorialUpdateController{UpdateHistorialUC: updateHistorialUC}
	historialDeleteController := &historialControllers.HistorialDeleteController{DeleteHistorialUC: deleteHistorialUC}

		historialConsumer, err := messaging.NewHistorialConsumer(createHistorialUC)
	if err != nil {
		log.Fatalf("Error al inicializar consumidor de historial: %v", err)
	}
go historialConsumer.StartConsuming()
	historialRoutes.HistorialsRoutes(router, historialCreateController, historialGetAllController, historialUpdateController, historialDeleteController, historialGetController)
}
