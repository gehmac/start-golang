package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAnimeForUserHandle(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Anime List",
	})
}
