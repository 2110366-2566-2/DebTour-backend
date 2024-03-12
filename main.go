package main

import (
	"DebTour/controllers"
	"DebTour/database"
	"DebTour/docs"
	"DebTour/middleware"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

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
	{

		v1.GET("/hello", controllers.HelloWorld)
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:username", controllers.GetUserByUsername)
		v1.POST("/users", controllers.CreateUser)
		v1.DELETE("/users/:username", controllers.DeleteUserByUsername)
		v1.PUT("/users/:username", controllers.UpdateUserByUsername)

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

		v1.GET("/agencies", controllers.GetAllAgencies)
		v1.GET("/agencies/:username", controllers.GetAgencyByUsername)
		v1.POST("/agencies", controllers.RegisterAgency)
		v1.PUT("/agencies/:username", controllers.UpdateAgency)
		v1.DELETE("/agencies/:username", controllers.DeleteAgency)

		v1.GET("/tourists", controllers.GetAllTourists)
		v1.GET("/tourists/:username", controllers.GetTouristByUsername)
		v1.POST("/tourists", controllers.RegisterTourist)
		//v1.POST("/tourists", controllers.CreateTourist)
		v1.PUT("/tourists/:username", controllers.UpdateTouristByUsername)
		v1.DELETE("/tourists/:username", controllers.DeleteTouristByUsername)

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

		v1.GET("/google/login", controllers.HandleGoogleLogin)
		v1.GET("/google/callback", controllers.HandleGoogleCallback)
		v1.GET("/google/logout", controllers.HandleGoogleLogout)
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
		v1.GET("/validatetoken/:token", controllers.ValidateTokenHandler)
		v1.GET("/validaterole/:token", controllers.ValidateRoleHandler)

		v1.GET("/issues", controllers.GetAllIssues)
		v1.POST("/issues", controllers.CreateIssueReport)

		v1.GET("/testdir3", controllers.TestRedir)
		v1.GET("/testdir2", controllers.TestDir)
		v1.GET("/google/testlogin/login", controllers.TestLogin)
		v1.GET("/google/testlogin/register", controllers.TestRegister)
	}

	v2 := router.Group("/api/v2")
	{
		v2.Use(middleware.AuthorizeJWT([]string{"admin", "tourist"}))
		v2.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	// a: admin, g: agency, t: tourist
	v2_a := router.Group("/api/v2/admin")
	{
		v2_a.Use(middleware.AuthorizeJWT([]string{"admin"}))
		v2_a.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	v2_ag := router.Group("/api/v2/agency")
	{
		v2_ag.Use(middleware.AuthorizeJWT([]string{"admin", "agency"}))
		v2_ag.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	v2_at := router.Group("/api/v2/tourist")
	{
		v2_at.Use(middleware.AuthorizeJWT([]string{"admin", "tourist"}))
		v2_at.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	v2_agt := router.Group("/api/v2/all")
	{
		v2_agt.Use(middleware.AuthorizeJWT([]string{"admin", "agency", "tourist"}))
		v2_agt.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	err := router.Run(":9000")
	if err != nil {
		return
	}

}
