package router

import "github.com/gin-gonic/gin"

type relation struct {
}

func (relation) Init(router *gin.RouterGroup) {
	r := router.Group("relation")
	{
		r.POST("action")
		r.GET("follow/list")
		r.GET("follower/list")
		r.GET("friend/list")
	}
}
