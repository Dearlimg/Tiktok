package router

import "github.com/gin-gonic/gin"

type comment struct {
}

func (c *comment) Init(router *gin.RouterGroup) {
	r := router.Group("comment")
	{
		r.POST("action")
		r.GET("list")
	}
}
