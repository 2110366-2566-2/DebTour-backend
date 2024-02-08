package main

import (
	"DebTour/controllers"
	"DebTour/models"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"DebTour/docs"
)

func SetUpSwagger() {
	docs.SwaggerInfo.Title = "DebTour API"
	docs.SwaggerInfo.Description = "DebTour API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func main() {

	models.InitDB()

	router := SetupRouter()

	SetUpSwagger()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", controllers.HelloWorld)
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:username", controllers.GetUserByUsername)
		v1.POST("/users", controllers.CreateUser)
		v1.DELETE("/users/:username", controllers.DeleteUser)
		v1.PUT("/users", controllers.UpdateUser)

		v1.GET("/tours", controllers.GetAllTours)
		v1.POST("/tours", controllers.CreateTour)
		v1.PUT("/tours", controllers.UpdateTour)
		v1.DELETE("/tours/:id", controllers.DeleteTour)
	}

	router.Run(":9000")
}
