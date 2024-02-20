package main

import (
	"DebTour/controllers"
	"DebTour/models"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"DebTour/docs"

	"github.com/zalando/gin-oauth2/google"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpSwagger() {
	docs.SwaggerInfo.Title = "DebTour API"
	docs.SwaggerInfo.Description = "DebTour API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var sessionName string

func main() {
	models.InitDB()
	sessionName = controllers.InitOauth()

	router := SetupRouter()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(corsConfig))

	router.LoadHTMLFiles("index.html")
	SetUpSwagger()

	router.Use(google.Session(sessionName)) // new

	// SetupOauth()
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
		v1.GET("/tours/:id", controllers.GetTourByID)
		v1.GET("/tours/tourists/:id", controllers.GetTouristByTourId)
		v1.POST("/tours", controllers.CreateTour)
		v1.PUT("/tours/:id", controllers.UpdateTour)
		v1.DELETE("/tours/:id", controllers.DeleteTour)
		v1.GET("/tours/filter", controllers.FilterTours)

		v1.GET("/activities", controllers.GetAllActivities)

		v1.POST("/joinings", controllers.JoinTour)
		v1.GET("/joinings", controllers.GetAllJoinings)

		v1.GET("/login", google.LoginHandler)
		v1.Use(google.Auth())
		v1.GET("/auth", controllers.UserInfoHandler)
		v1.GET("/api", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello from private for groups"})
		})
		v1.GET("/logout", controllers.OauthLogout)
	}

	router.Run(":9000")
}
