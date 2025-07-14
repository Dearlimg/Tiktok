package router

import "github.com/gin-gonic/gin"

type message struct {
}

func (msg *message) Init(router *gin.RouterGroup) {}
