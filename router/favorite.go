package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type favorite struct {
}

func (favorite *favorite) Init(router *gin.RouterGroup) {
	r := router.Group("/favorite")
	{
		r.GET("list", middleware.AuthWithoutLogin(), api.FavoriteList)
		r.POST("action", middleware.Auth(), api.FavoriteAction)
	}
}
