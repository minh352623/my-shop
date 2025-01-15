package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			name := c.Query("name")
			age := c.DefaultQuery("age", "20")
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"name":    name,
				"age":     age,
			})
		})
	}
	return r
}
