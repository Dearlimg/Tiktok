package router

import "github.com/gin-gonic/gin"

type message struct {
}

func (msg *message) Init(router *gin.RouterGroup) {
	r := router.Group("/message")
	{
		r.GET("chat")
		r.POST("action")
	}
}
