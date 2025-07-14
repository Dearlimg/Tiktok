package router

import "github.com/gin-gonic/gin"

type publish struct {
}

func (p *publish) Init(router *gin.RouterGroup) {
	r := router.Group("publish")
	{
		r.GET("list")
		r.POST("action")
	}
}
