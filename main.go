package main

import (
	"ppap/backup/go/config"
	"ppap/backup/go/models"
	"ppap/backup/go/routers"
)

func main() {
	// config 模块必须得最先初始化
	config.Setup("./config.ini")
	// 初始化数据库层
	models.Setup()

	engine := routers.InitRouter()
	engine.Run(":8080")
}

