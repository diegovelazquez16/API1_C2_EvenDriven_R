package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"holamundo/users/aplication/usecase"
	"holamundo/users/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateController struct {
	CreateUserUC *usecase.CreateUserUseCase
}

func (c *UserCreateController) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}

	user.Password = string(hashedPassword)

	err = c.CreateUserUC.Execute(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado de forma exitosa",
		"user":    user,
	})
}
