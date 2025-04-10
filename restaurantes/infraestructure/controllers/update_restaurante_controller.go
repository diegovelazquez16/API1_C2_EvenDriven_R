package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/restaurantes/aplication/usecase"
	"holamundo/restaurantes/domain/models"
)

type RestauranteUpdateController struct {
	UpdateRestauranteUC *usecase.UpdateRestauranteUseCase
}

func (c *RestauranteUpdateController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Restaurante ID",
		})
		return
	}

	var restaurante models.Restaurante
	if err := ctx.ShouldBindJSON(&restaurante); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	restaurante.ID = uint(id)

	if err := c.UpdateRestauranteUC.Execute(&restaurante); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fallo al actualizar al usuario: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Restaurante actualizado exitosamente",
		"restauranrte":    restaurante,
	})
}