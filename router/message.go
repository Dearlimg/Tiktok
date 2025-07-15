package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type message struct {
}

func (msg *message) Init(router *gin.RouterGroup) {
	r := router.Group("/message")
	{
		r.GET("chat", middleware.Auth(), api.MessageChat)
		r.POST("action", middleware.Auth(), api.MessageAction)
	}
}
