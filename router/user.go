package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type User struct {
}

func (u User) Init(router *gin.RouterGroup) {
	r := router.Group("user")
	{
		userGroup := r.Group("")
		{
			userGroup.GET("", middleware.Auth(), api.UserInfo)
			userGroup.POST("register", api.Register)
			userGroup.POST("login", api.Login)
		}
	}
}
