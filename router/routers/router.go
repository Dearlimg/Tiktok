package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok001/router"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	root := r.Group("")
	{
		root.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		rg := router.Routers
		rg.User.Init(root)
		rg.Message.Init(root)
		rg.Relation.Init(root)
		rg.Comment.Init(root)
		rg.Feed.Init(root)
		rg.Favorite.Init(root)
		rg.Publish.Init(root)
	}
	return r
}
