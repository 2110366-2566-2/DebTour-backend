package main

import (
	"DebTour/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func main() {
	router := SetupRouter()

	router.GET("/", controllers.HelloWorld)

	router.Run("localhost:9000")
}
