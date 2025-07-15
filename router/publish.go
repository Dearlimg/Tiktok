package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type publish struct {
}

func (p *publish) Init(router *gin.RouterGroup) {
	r := router.Group("publish")
	{
		r.GET("list", middleware.AuthWithoutLogin(), api.PublishList)
		r.POST("action", middleware.AuthBody(), api.Publish)
	}
}
