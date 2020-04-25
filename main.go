package main

import "ppap/backup/go/routers"

func main() {
	engine := routers.InitRouter()
	engine.Run(":8080")
}
