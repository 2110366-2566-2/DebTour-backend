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

		v1.GET("/hello", controllers.HelloWorld)                  // all
		v1.GET("/users", controllers.GetAllUsers)                 // admin
		v1.GET("/users/:username", controllers.GetUserByUsername) // admin
		v1.GET("/getMe", controllers.GetMe)                       // logged in

		v1.GET("/tours", controllers.GetAllTours)                          // all
		v1.GET("/tours/:id", controllers.GetTourByID)                      // all
		v1.GET("/tours/tourists/:id", controllers.GetTouristByTourId)      // admin, agency owner
		v1.POST("/tours", controllers.CreateTour)                          // agency
		v1.PUT("/tours/:id", controllers.UpdateTour)                       // admin, agency owner
		v1.DELETE("/tours/:id", controllers.DeleteTour)                    // admin, agency owner
		v1.GET("/tours/filter", controllers.FilterTours)                   // all
		v1.PUT("/tours/activities/:id", controllers.UpdateTourActivities)  // admin, agency owner
		v1.POST("/tours/activities/:id", controllers.CreateTourActivities) // agency owner

		v1.GET("/tours/images/:id", controllers.GetTourImages)               // all
		v1.POST("/tours/images/:id", controllers.CreateTourImagesByTourId)   // agency owner
		v1.DELETE("/tours/images/:id", controllers.DeleteTourImagesByTourId) // admin, agency owner

		//admin
		v1.GET("/agencies", controllers.GetAllAgencies)
		v1.GET("/agencies/:username", controllers.GetAgencyByUsername)
		v1.PUT("/agencies/:username", controllers.UpdateAgency) // + agency themselves
		v1.DELETE("/agencies/:username", controllers.DeleteAgency)
		//end admin
		// getme() // logged in

		v1.GET("/tourists", controllers.GetAllTourists)                       // admin
		v1.GET("/tourists/:username", controllers.GetTouristByUsername)       // all + login
		v1.PUT("/tourists/:username", controllers.UpdateTouristByUsername)    // admin, tourist themselves
		v1.DELETE("/tourists/:username", controllers.DeleteTouristByUsername) //admin

		v1.GET("/activities", controllers.GetAllActivities) //admin

		v1.POST("/joinings", controllers.JoinTour)      // tourist
		v1.GET("/joinings", controllers.GetAllJoinings) // admin

		v1.GET("/reviews", controllers.GetAllReviews)               // admin
		v1.GET("/reviews/:id", controllers.GetReviewById)           // all
		v1.GET("/reviews/tour/:id", controllers.GetReviewsByTourId) // all
		v1.POST("/reviews/tour/:id", controllers.CreateReview)      // tourist
		v1.GET("/reviews/tourist/:username", controllers.GetReviewsByTouristUsername)
		v1.DELETE("/reviews/:id", controllers.DeleteReview)                                 // admin
		v1.DELETE("/reviews/tour/:id", controllers.DeleteReviewsByTourId)                   // admin, agency
		v1.DELETE("/reviews/tourist/:username", controllers.DeleteReviewsByTouristUsername) // admin, tourist
		v1.GET("/reviews/averageRating/:id", controllers.GetAverageRatingByTourId)          // all

		// auth
		v1.POST("/auth/registerTourist", controllers.RegisterTourist)
		v1.POST("/auth/registerAgency", controllers.RegisterAgency)
		v1.POST("/auth/login", controllers.Login)
		v1.GET("/auth/logout", controllers.HandleGoogleLogout)
		//end auth

		v1.GET("/issues", controllers.GetIssues)                   // all + logged in and only allowed (addmin = all, tourist+agency = only their own)
		v1.POST("/issues", controllers.CreateIssueReport)          // all + logged in
		v1.PUT("/issues/:issue_id", controllers.UpdateIssueReport) // admin
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
		v2_ag.Use(middleware.AuthorizeJWT([]string{"admin", "Agency"}))
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
