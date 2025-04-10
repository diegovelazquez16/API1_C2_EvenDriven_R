package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/historial/aplication/usecase"
)

type HistorialGetController struct {
	GetHistorialUC *usecase.GetHistorialsUseCase
}

func (c *HistorialGetController) GetHistorialByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe estar registrado o ser un numero valido"})
		return
	}

	Historial, err := c.GetHistorialUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, Historial)
}
