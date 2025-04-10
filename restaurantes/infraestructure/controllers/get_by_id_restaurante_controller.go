package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/restaurantes/aplication/usecase"
)

type RestauranteGetController struct {
	GetRestauranteUC *usecase.GetRestauranteUseCase
}

func (c *RestauranteGetController) GetRestauranteByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe estar registrado o ser un numero valido"})
		return
	}

	restaurante, err := c.GetRestauranteUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, restaurante)
}
