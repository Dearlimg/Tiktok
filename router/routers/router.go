package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok001/router"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	root := r.Group("tiktok")
	{
		root.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		rg := router.Routers
		rg.User.Init(root)     //finished
		rg.Message.Init(root)  //finished
		rg.Relation.Init(root) //finished
		rg.Comment.Init(root)  //finished
		rg.Feed.Init(root)     //finished
		rg.Favorite.Init(root) //finished
		rg.Publish.Init(root)  //finished
	}
	return r
}
