package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"holamundo/users/aplication/usecase"
	"holamundo/users/domain/models"
)

type UserUpdateController struct {
	UpdateUserUC *usecase.UpdateUserUseCase
}

func (c *UserUpdateController) Update(ctx *gin.Context) {
	// Obtener el ID de la URL y validarlo
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// Mapear el JSON del cuerpo de la solicitud a la estructura User
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	// Asignar el ID de la URL al objeto User
	user.ID = uint(id)

	// Ejecutar el caso de uso para actualizar el usuario
	if err := c.UpdateUserUC.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user: " + err.Error(),
		})
		return
	}

	// Devolver una respuesta exitosa con el usuario actualizado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}