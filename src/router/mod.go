package router

import (
	"github.com/devproje/project-mirror/src/config"
	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine) {
	engine.GET("/", mirror)
	engine.GET("/:path/*child", mirrorPath)
	engine.Static("/public", "public")
	if config.Get().Auth {
		v1 := engine.Group("/v1")
		{
			v1.POST("/login", Login)
			v1.GET("/login", LoginForm)
		}
	}
}
