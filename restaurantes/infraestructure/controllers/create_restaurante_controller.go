package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"holamundo/restaurantes/aplication/usecase"
	"holamundo/restaurantes/domain/models"
)

type RestauranteCreateController struct {
	CreateRestauranteUC *usecase.CreateRestauranteUseCase
}	

func (c *RestauranteCreateController) Create(ctx *gin.Context) {
	var restaurante models.Restaurante
	if err := ctx.ShouldBindJSON(&restaurante); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreateRestauranteUC.Execute(&restaurante)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Restaurante creado de forma exitosa",
		"restaurante": restaurante,
	})
}
