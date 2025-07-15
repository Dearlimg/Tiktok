package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type feed struct {
}

func (f *feed) Init(router *gin.RouterGroup) {
	r := router.Group("")
	{
		r.GET("feed", middleware.AuthWithoutLogin(), api.Feed)
	}
}
