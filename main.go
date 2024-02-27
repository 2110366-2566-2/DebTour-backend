package main

import (
	"DebTour/controllers"
	"DebTour/database"
	"DebTour/middleware"
	"net/http"
	"os"

	//"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	// "DebTour/docs"
	// "DebTour/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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

func SetupOauth() {
	controllers.InitializeOauthenv()
}

func main() {
	database.InitDB()

	router := SetupRouter()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(corsConfig))

	SetUpSwagger()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Set up Oauth
	SetupOauth()

	v1 := router.Group("/api/v1")
	// v1.Use(middleware.AuthorizeJWT())
	{

		v1.GET("/hello", controllers.HelloWorld)
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:username", controllers.GetUserByUsername)
		v1.POST("/users", controllers.CreateUser)
		v1.DELETE("/users", controllers.DeleteUser)
		v1.PUT("/users", controllers.UpdateUser)

		v1.GET("/tours", controllers.GetAllTours)
		v1.GET("/tours/:id", controllers.GetTourByID)
		v1.GET("/tours/tourists/:id", controllers.GetTouristByTourId)
		v1.POST("/tours", controllers.CreateTour)
		v1.PUT("/tours/:id", controllers.UpdateTour)
		v1.DELETE("/tours/:id", controllers.DeleteTour)
		v1.GET("/tours/filter", controllers.FilterTours)
		v1.PUT("/tours/activities/:id", controllers.UpdateTourActivities)
		v1.POST("/tours/activities/:id", controllers.CreateTourActivities)

		v1.GET("/tours/images/:id", controllers.GetTourImages)
		v1.POST("/tours/images/:id", controllers.CreateTourImagesByTourId)
		v1.DELETE("/tours/images/:id", controllers.DeleteTourImagesByTourId)

		v1.GET("/activities", controllers.GetAllActivities)

		v1.POST("/joinings", controllers.JoinTour)
		v1.GET("/joinings", controllers.GetAllJoinings)

		v1.GET("/reviews", controllers.GetAllReviews)
		v1.GET("/reviews/:id", controllers.GetReviewById)
		v1.GET("/reviews/tour/:id", controllers.GetReviewsByTourId)
		v1.POST("/reviews/tour/:id", controllers.CreateReview)
		v1.GET("/reviews/tourist/:username", controllers.GetReviewsByTouristUsername)
		v1.DELETE("/reviews/:id", controllers.DeleteReview)
		v1.DELETE("/reviews/tour/:id", controllers.DeleteReviewsByTourId)
		v1.DELETE("/reviews/tourist/:username", controllers.DeleteReviewsByTouristUsername)

		v1.GET("/google/login/:role", controllers.HandleGoogleLogin)
		v1.GET("/google/callback", controllers.HandleGoogleCallback)
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
		v1.GET("/validatetoken/:token", controllers.ValidateTokenHandler)
		v1.GET("/validaterole/:token", controllers.ValidateRoleHandler)

		v1.GET("/issues", controllers.GetAllIssues)
	}
	v2 := router.Group("/api/v2")
	v2.Use(middleware.AuthorizeJWT())
	{
		v2.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
	err := router.Run(":9000")
	if err != nil {
		return
	}

}
