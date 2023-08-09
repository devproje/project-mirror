package router

import "github.com/gin-gonic/gin"

func New(engine *gin.Engine) {
	engine.GET("/", mirror)
	engine.GET("/:path/*child", mirrorPath)
	engine.Static("/file", ".data")
	engine.Static("/static", "static")
}
