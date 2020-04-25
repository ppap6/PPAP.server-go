package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	engine := gin.Default()
	return engine
}
