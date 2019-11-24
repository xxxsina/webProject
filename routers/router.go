package routers

import (
	"github.com/gin-gonic/gin"
	"webProject/controllers"
)

func RegisterRouters(g *gin.Engine) {
	g.GET("/name/:id", controllers.GetName)
	v1 := g.Group("/v1")
	{
		v1.POST("/add", controllers.Add)
	}
}