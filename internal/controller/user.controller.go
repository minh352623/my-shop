package controller

import (
	"ecom/internal/service"
	"ecom/pkg/response"

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
	response.SuccessResponse(ctx, response.Success, gin.H{"id": id, "name": name})
}
