package router

import (
	"job-openings/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	handler.InitHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpening)
		v1.POST("/opening", handler.CreateOpening)
		v1.PUT("/opening", handler.UpdateOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.GET("/openings", handler.ListOpenings)
	}
}
