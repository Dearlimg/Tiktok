package router

import (
	"github.com/gin-gonic/gin"
	"tiktok001/controller/api"
	"tiktok001/middleware"
)

type relation struct {
}

func (relation) Init(router *gin.RouterGroup) {
	r := router.Group("relation")
	{
		r.POST("action", middleware.Auth(), api.RelationAction)
		r.GET("follow/list", middleware.Auth(), api.FollowList)
		r.GET("follower/list", middleware.Auth(), api.FollowerList)
		r.GET("friend/list", middleware.Auth(), api.FriendList)
	}
}
