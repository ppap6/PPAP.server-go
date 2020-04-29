package routers

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"ppap/backup/go/logger"
)

func InitRouter() *gin.Engine{

	// gin 日志单独记录
	f, err := os.OpenFile("./logs/gin.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.Error(nil, "init gin log file failed", "err", err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	engine := gin.Default()
	engine.Use(logger.RequestID())


	engine.GET("ping", func(c *gin.Context) {
		logger.Info(c, "ping", "response", "pong")
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
