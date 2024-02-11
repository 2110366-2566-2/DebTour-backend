package main

import (
	"DebTour/controllers"
	"DebTour/models"

	"github.com/gin-gonic/gin"

	"DebTour/docs"

	// "os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/gorilla/pat"
	//"github.com/gorilla/sessions"
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

func SetupOauth() {
	controllers.InitializeOauthenv()
}

func main() {
	models.InitDB()

	router := SetupRouter()
	router.LoadHTMLFiles("index.html")
	SetUpSwagger()
	SetupOauth()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/", controllers.HandleMain)
	// router.GET("/GoogleLogin", controllers.HandleGoogleLogin)
	// router.GET("/GoogleCallback", controllers.HandleGoogleCallback)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", controllers.HelloWorld)
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:username", controllers.GetUserByUsername)
		v1.POST("/users", controllers.CreateUser)
		v1.DELETE("/users/:username", controllers.DeleteUser)
		v1.PUT("/users", controllers.UpdateUser)

		v1.GET("/tours", controllers.GetAllTours)
		v1.GET("/tours/:id", controllers.GetTourByID)
		v1.POST("/tours", controllers.CreateTour)
		v1.PUT("/tours/:id", controllers.UpdateTour)
		v1.DELETE("/tours/:id", controllers.DeleteTour)
		v1.GET("/tours/filter", controllers.FilterTours)

		v1.POST("/joinings", controllers.JoinTour)
		v1.GET("/joinings", controllers.GetAllJoinings)
		// v1.GET("/", controllers.HandleMain)
		// v1.GET("/GoogleLogin", controllers.HandleGoogleLogin)
		// v1.GET("/GoogleCallback", controllers.HandleGoogleCallback)
		v1.GET("/", controllers.HandleMain)
		v1.GET("/GoogleLogin", controllers.HandleGoogleLogin)
		v1.GET("/GoogleCallback", controllers.HandleGoogleCallback)
	}

	router.Run(":9000")
}
