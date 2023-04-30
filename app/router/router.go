package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	setupUserRoutes(r)
	setupPostRoutes(r)
}