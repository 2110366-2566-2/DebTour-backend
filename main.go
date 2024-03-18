package main

import (
	"DebTour/controllers"
	"DebTour/database"
	"DebTour/docs"
	"DebTour/middleware"
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

var ()

// @securityDefinitions.apiKey ApiKeyAuth
// @in       header
// @name     Authorization
// @description Type "Bearer " followed by a space and then your token
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
		v1.GET("/givemetoken/:username", controllers.GetToken)
		v1.GET("givemeusername/:token", controllers.GetUsername)
		v1.GET("/logout", middleware.AuthorizeJWT([]string{"Agency", "Tourist"}), controllers.HandleGoogleLogout)

		v1.GET("/hello", controllers.HelloWorld)                                                              // all
		v1.GET("/users", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllUsers)                 // admin
		v1.GET("/users/:username", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetUserByUsername) // admin
		v1.GET("/getMe", middleware.AuthorizeJWT([]string{"Agency", "Tourist"}), controllers.GetMe)           // logged in

		v1.GET("/tours", controllers.GetAllTours)                                                                                  // all
		v1.GET("/tours/:id", controllers.GetTourByID)                                                                              // all
		v1.GET("/tours/tourists/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2), controllers.GetTouristByTourId)     // admin, agency owner
		v1.POST("/tours", middleware.AuthorizeJWT([]string{"Agency"}), controllers.CreateTour)                                     // agency
		v1.PUT("/tours/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2), controllers.UpdateTour)                      // admin, agency owner
		v1.DELETE("/tours/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2), controllers.DeleteTour)                   // admin, agency owner
		v1.GET("/tours/filter", controllers.FilterTours)                                                                           // all
		v1.PUT("/tours/activities/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2), controllers.UpdateTourActivities) // admin, agency owner
		v1.POST("/tours/activities/:id", middleware.AuthorizeJWT([]string{"Agency"}), controllers.CreateTourActivities)            // agency owner

		v1.GET("/tours/images/:id", controllers.GetTourImages)                                                                        // all
		v1.POST("/tours/images/:id", middleware.AuthorizeJWT([]string{"Agency"}), controllers.CreateTourImagesByTourId)               // agency owner
		v1.DELETE("/tours/images/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2), controllers.DeleteTourImagesByTourId) // admin, agency owner

		//admin
		v1.GET("/agencies", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllAgencies)
		v1.GET("/agencies/:username", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAgencyByUsername)
		v1.PUT("/agencies/:username", middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 1), controllers.UpdateAgency) // + agency themselves
		v1.DELETE("/agencies/:username", middleware.AuthorizeJWT([]string{"Admin"}), controllers.DeleteAgency)
		v1.GET("/agencies/companyInformation", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllAgenciesWithCompanyInformation)
		//end admin

		v1.GET("/tourists", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllTourists)                                      // admin
		v1.GET("/tourists/:username", middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}), controllers.GetTouristByUsername) // all + login
		v1.PUT("/tourists/:username", middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1), controllers.UpdateTouristByUsername)     // admin, tourist themselves
		v1.DELETE("/tourists/:username", middleware.AuthorizeJWT([]string{"Admin"}), controllers.DeleteTouristByUsername)                // admin

		v1.GET("/activities", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllActivities) //admin

		v1.POST("/joinings", middleware.AuthorizeJWT([]string{"Tourist"}), controllers.JoinTour)    // tourist
		v1.GET("/joinings", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllJoinings) // admin

		v1.GET("/reviews", middleware.AuthorizeJWT([]string{"Admin"}), controllers.GetAllReviews)                                                  // admin
		v1.GET("/reviews/:id", controllers.GetReviewById)                                                                                          // all
		v1.GET("/reviews/tour/:id", controllers.GetReviewsByTourId)                                                                                // all
		v1.POST("/reviews/tour/:id", middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}), controllers.CreateReview)                    // tourist
		v1.GET("/reviews/tourist/:username", middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1), controllers.GetReviewsByTouristUsername)    // admin, tourist themselves
		v1.DELETE("/reviews/:id", middleware.AuthorizeJWT([]string{"Admin"}), controllers.DeleteReview)                                            // admin
		v1.DELETE("/reviews/tour/:id", middleware.AuthorizeJWT([]string{"Admin", "Agency"}), controllers.DeleteReviewsByTourId)                    // admin, agency
		v1.DELETE("/reviews/tourist/:username", middleware.AuthorizeJWT([]string{"Admin", "Tourist"}), controllers.DeleteReviewsByTouristUsername) // admin, tourist
		v1.GET("/reviews/averageRating/:id", controllers.GetAverageRatingByTourId)                                                                 // all

		// auth
		v1.POST("/auth/registerTourist", controllers.RegisterTourist)
		v1.POST("/auth/registerAgency", controllers.RegisterAgency)
		v1.POST("/auth/login", controllers.Login)
		v1.GET("/auth/logout", controllers.HandleGoogleLogout)
		//end auth

		v1.GET("/issues", middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}), controllers.GetIssues)          // all + logged in and only allowed (addmin = all, tourist+agency = only their own)
		v1.POST("/issues", middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}), controllers.CreateIssueReport) // all + logged in
		v1.PUT("/issues/:issue_id", middleware.AuthorizeJWT([]string{"Admin"}), controllers.UpdateIssueReport)             // admin

	}

	err := router.Run(":9000")
	if err != nil {
		return
	}

}
