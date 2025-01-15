package controller

import (
	"ecom/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, name := uc.userService.GetUser(ctx.Param("id"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    gin.H{"id": id, "name": name},
	})
}
