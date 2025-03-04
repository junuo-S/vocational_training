package main

import (
	"webservice/src/global"

	"github.com/gin-gonic/gin"
)

func main() {
	rootRouter := gin.New()
	rootRouter.Use(global.JunuoWebLogMiddleware(), global.AntiSpiderMiddleware(), gin.Recovery())
	rootRouter.Run(":3000")
}
