package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type comment struct {
}

func (c *comment) Init(router *gin.RouterGroup) {
	r := router.Group("comment")
	{
		r.POST("action", middleware.Auth(), api.CommentAction)
		r.GET("list", middleware.AuthWithoutLogin(), api.CommentList)
	}
}
