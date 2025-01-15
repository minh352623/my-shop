package routers

import (
	c "ecom/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/api")
	{
		v1.GET("/users/:id", c.NewUserController().GetUser)
	}
	return r
}
