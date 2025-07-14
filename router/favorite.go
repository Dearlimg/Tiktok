package router

import "github.com/gin-gonic/gin"

type favorite struct {
}

func (favorite *favorite) Init(router *gin.RouterGroup) {
	r := router.Group("/favorite")
	{
		r.GET("list")
		r.POST("action")
	}
}
