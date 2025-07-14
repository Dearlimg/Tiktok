package router

import "github.com/gin-gonic/gin"

type feed struct {
}

func (f *feed) Init(router *gin.RouterGroup) {
	r := router.Group("")
	{
		r.GET("feed")
	}

}
