package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"holamundo/restaurantes/aplication/usecase"
)
type RestauranteGetAllController struct {
	GetAllRestaurantesUC *usecase.GetAllRestaurantesUseCase
}
//cambia GetAll a execut o run y lo mismo para todos
func (c *RestauranteGetAllController) GetAll(ctx *gin.Context) {
	restaurantes, err := c.GetAllRestaurantesUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, restaurantes)
}