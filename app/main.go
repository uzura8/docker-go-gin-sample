package main

import (
	"gin-sample/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8000")
}