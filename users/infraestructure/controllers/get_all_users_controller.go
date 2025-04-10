package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"holamundo/users/aplication/usecase"
)
type UserGetAllController struct {
	GetAllUsersUC *usecase.GetAllUsersUseCase
}
//cambia GetAll a execut o run y lo mismo para todos
func (c *UserGetAllController) GetAll(ctx *gin.Context) {
	users, err := c.GetAllUsersUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}