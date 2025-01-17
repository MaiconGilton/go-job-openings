package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	InitRoutes(router)

	router.Run()
}
