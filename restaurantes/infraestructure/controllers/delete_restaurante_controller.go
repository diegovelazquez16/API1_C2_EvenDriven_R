package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"holamundo/restaurantes/aplication/usecase"
)

type RestauranteDeleteController struct {
	DeleteRestauranteUC *usecase.DeleteRestauranteUseCase
}

func (c *RestauranteDeleteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inavlido"})
		return
	}

	err = c.DeleteRestauranteUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Restaurante eliminado"})
}