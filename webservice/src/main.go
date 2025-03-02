package main

import (
	"webservice/src/global"

	"github.com/gin-gonic/gin"
)

func main() {
	rootRouter := gin.New()
	rootRouter.Use(global.JunuoWebLog(), gin.Recovery())
	rootRouter.Run(":3000")
}
