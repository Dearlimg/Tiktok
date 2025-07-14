package router

import "github.com/gin-gonic/gin"

type User struct {
}

func (u User) Init(router *gin.RouterGroup) {
	r := router.Group("user")
	{
		userGroup := r.Group("")
		{
			userGroup.GET("")
			userGroup.POST("register")
			userGroup.POST("login")
		}
	}
}
