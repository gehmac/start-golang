package router

import (
	"github.com/gehmac/anime-list-golang/handler"
	"github.com/gin-gonic/gin"
)

func startRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("my-list", handler.ListAnimeForUserHandle)
	}
}
