package router

import (
	"gin-sample/handlers"

	"github.com/gin-gonic/gin"
)

func setupPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/posts")
	{
		postGroup.GET("/", handlers.GetPosts)
		postGroup.GET("/:id", handlers.GetPost)
		postGroup.POST("/", handlers.CreatePost)
		postGroup.PUT("/:id", handlers.UpdatePost)
		postGroup.DELETE("/:id", handlers.DeletePost)
	}
}