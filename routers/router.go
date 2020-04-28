package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	engine := gin.Default()

	engine.GET("ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	engine.GET("pong", func(c *gin.Context) {
		c.JSON(200, "ping")
	})
	engine.GET("ping/pong", func(c *gin.Context) {
		c.JSON(200, "pong/ping")
	})

	return engine
}
