package router

import (
	"gin-sample/handlers"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.POST("/", handlers.CreateUser)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}
}