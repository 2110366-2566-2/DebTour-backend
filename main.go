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
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
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
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	router.Use(cors.New(corsConfig))

	SetUpSwagger()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		// Debugging purposes
		v1.GET("/givemetoken/:username", controllers.GetToken)
		v1.GET("givemeusername/:token", controllers.GetUsername)
		v1.GET("/hello", controllers.HelloWorld)

		// Users
		// Get all users (admin)
		v1.GET("/users",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllUsers)
		// Delete user by username (admin)
		v1.DELETE("/users/:username",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteUserByUsername)
		// Get user by username (admin)
		v1.GET("/users/:username",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetUserByUsername)
		// Get information about the user that is logged in (all)
		v1.GET("/getMe",
			middleware.AuthorizeJWT([]string{"Admin", "Agency", "Tourist"}),
			controllers.GetMe)

		// Tours
		// Get all tours (all)
		v1.GET("/tours", controllers.GetAllTours)
		// Get tour by id (all)
		v1.GET("/tours/:id", controllers.GetTourByID)
		// Create tour along with tour images and tour activities (agency)
		v1.POST("/tours",
			middleware.AuthorizeJWT([]string{"Agency"}),
			controllers.CreateTour)
		// Update tour along with tour images (admin, agency owner)
		v1.PUT("/tours/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.UpdateTour)
		// Delete tour by id (admin, agency owner)
		v1.DELETE("/tours/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.DeleteTour)
		// Get filtered tours (all)
		v1.GET("/tours/filter", controllers.FilterTours)
		// Update tour activities (admin, agency owner)
		v1.PUT("/tours/activities/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.UpdateTourActivities)
		// Create tour activities (agency owner)
		v1.POST("/tours/activities/:id",
			middleware.AuthorizeJWT([]string{"Agency"}),
			controllers.CreateTourActivities)
		// Get all tours by agency username (admin, agency owner)
		v1.GET("/tours/agency/:username", controllers.GetToursByAgencyUsername)

		// Tour images
		// Get tour images (all)
		v1.GET("/tours/images/:id", controllers.GetTourImages)
		// Create tour images (agency owner)
		v1.POST("/tours/images/:id",
			middleware.AuthorizeJWT([]string{"Agency"}),
			controllers.CreateTourImagesByTourId)
		// Delete tour images by Tour id (admin, agency owner)
		v1.DELETE("/tours/images/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.DeleteTourImagesByTourId)

		// Agencies
		// Get all agencies (admin)
		v1.GET("/agencies",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllAgencies)
		// Get agency by username (admin)
		v1.GET("/agencies/:username",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAgencyWithUserByUsername)
		// Update agency by username (admin, agency owner)
		v1.PUT("/agencies/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 1),
			controllers.UpdateAgencyByUsername)

		// Company Information
		// Get all agencies with company information (admin)
		v1.GET("/agencies/companyInformation",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllAgenciesWithCompanyInformation)
		// Get company information by agency username (admin, agency owner)
		v1.GET("/agencies/companyInformation/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 1),
			controllers.GetCompanyInformationByAgencyUsername)
		// Delete company information by agency username (admin, agency owner)
		v1.DELETE("/agencies/companyInformation/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 1),
			controllers.DeleteCompanyInformationByAgencyUsername)
		// Verify agency (admin)
		v1.PUT("agencies/verify",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.VerifyAgency)

		// Agency Revenue
		// Get remaining revenue (agency owner)
		v1.GET("/agencies/getRevenue",
			middleware.AuthorizeJWT([]string{"Agency"}),
			controllers.GetRemainingRevenue)

		// Tourists
		// Get all tourists (admin)
		v1.GET("/tourists",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllTouristsWithUser)
		// Get tourist by username (all)
		v1.GET("/tourists/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}),
			controllers.GetTouristByUsername)
		// Update tourist by username (admin, tourist themselves)
		v1.PUT("/tourists/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1),
			controllers.UpdateTouristByUsername)

		// Activities
		// Get all activities (admin)
		v1.GET("/activities",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllActivities)

		// Joinings
		// Join tour (tourist)
		v1.POST("/joinings",
			middleware.AuthorizeJWT([]string{"Tourist"}),
			controllers.JoinTour)
		// Get all joinings (admin)
		v1.GET("/joinings",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllJoinings)
		// Get tourists by tour id (admin, agency owner)
		v1.GET("/tours/tourists/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.GetTouristByTourId)

		// Reviews
		// Get all reviews (admin)
		v1.GET("/reviews",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllReviews) // admin
		// Get review by id (all)
		v1.GET("/reviews/:id", controllers.GetReviewById) // all
		// Get reviews by tour id (all)
		v1.GET("/reviews/tour/:id", controllers.GetReviewsByTourId) // all
		// Create review (tourist)
		v1.POST("/reviews/tour/:id",
			middleware.AuthorizeJWT([]string{"Tourist"}),
			controllers.CreateReview) // tourist
		// Get reviews by tourist username (admin, tourist themselves)
		v1.GET("/reviews/tourist/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1),
			controllers.GetReviewsByTouristUsername) // admin, tourist themselves
		// Delete review by id (admin)
		v1.DELETE("/reviews/:id",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteReview) // admin
		// Delete reviews by tour id (admin, agency themselves)
		v1.DELETE("/reviews/tour/:id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.DeleteReviewsByTourId) // admin, agency themselves
		// Delete reviews by tourist username (admin, tourist themselves)
		v1.DELETE("/reviews/tourist/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1),
			controllers.DeleteReviewsByTouristUsername) // admin, tourist themselves
		// Get average rating by tour id (all)
		v1.GET("/reviews/averageRating/:id", controllers.GetAverageRatingByTourId) // all

		// Auth
		// Register tourist (all)
		v1.POST("/auth/registerTourist", controllers.RegisterTourist)
		// Register agency (all)
		v1.POST("/auth/registerAgency", controllers.RegisterAgency)
		// Login (all)
		v1.POST("/auth/login", controllers.Login)
		// Logout (all)
		v1.GET("/auth/logout",
			middleware.AuthorizeJWT([]string{"Admin", "Agency", "Tourist"}),
			controllers.HandleGoogleLogout)

		// Issues
		// Get all issues (admin)
		v1.GET("/issues",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}),
			controllers.GetIssues)
		// Create issue report (all)
		v1.POST("/issues",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist", "Agency"}),
			controllers.CreateIssueReport)
		// Update issue report (admin)
		v1.PUT("/issues/:issue_id",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.UpdateIssueReport)

		// TransactionPayments
		// Get all transaction payments (admin)
		v1.GET("/transactionPayments",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllTransactionPayments)
		// Get transaction payment by transaction id (all)
		v1.GET("/transactionPayments/:transactionId",
			middleware.AuthorizeJWT([]string{"Admin", "Agency", "Tourist"}),
			controllers.GetTransactionPaymentByTransactionId)
		// Get transaction payment by tour id (Admin, Agency Owner)
		v1.GET("/transactionPayments/tours/:tourId",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}, 2),
			controllers.GetTransactionPaymentByTourId)
		// Get transaction payment by tourist username (Admin, Tourist Owner)
		v1.GET("/transactionPayments/tourists/:username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}, 1),
			controllers.GetTransactionPaymentByTouristUsername)
		// Get Stripe public key (all)
		v1.GET("/stripePublicKey", controllers.GetStripePublicKey)
		// Start payment (tourist)
		v1.POST("/transactionPayments",
			middleware.AuthorizeJWT([]string{"Tourist"}),
			controllers.StartTransactionPayment)
		// Update transaction status (tourist and admin)
		v1.PUT("/transactionPayments/:transactionId",
			middleware.AuthorizeJWT([]string{"Tourist", "Admin"}),
			controllers.UpdateTransactionStatus)
		// refund transaction by transaction id (admin)
		v1.PUT("/transactionPayments/refund/:transactionId",
			middleware.AuthorizeJWT([]string{"Tourist"}),
			controllers.RefundTransaction)
		// Delete transaction payment by transaction id (admin)
		v1.DELETE("/transactionPayments/:transactionId",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteTransactionPayment)
		// Delete transaction payment by tour id (admin)
		v1.DELETE("/transactionPayments/tours/:tourId",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteTransactionPaymentByTourId)
		// Delete transaction payment by tourist username (admin)
		v1.DELETE("/transactionPayments/tourists/:username",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteTransactionPaymentByTouristUsername)

		// Suggestion
		// Get all suggestions (admin)
		v1.GET("/suggestions",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetAllSuggestions)
		// Get suggestion by suggestion id (admin)
		v1.GET("/suggestions/:suggestion_id",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.GetSuggestionBySuggestionId)
		// Create suggestion (tourist)
		v1.POST("/suggestions",
			middleware.AuthorizeJWT([]string{"Tourist"}),
			controllers.CreateSuggestion)
		// Get all suggestions with location (admin, agency)
		v1.GET("/suggestions/location",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}),
			controllers.GetAllSuggestionsWithLocation)
		// Get suggestion with location by suggestion id (admin, agency)
		v1.GET("/suggestions/location/:suggestion_id",
			middleware.AuthorizeJWT([]string{"Admin", "Agency"}),
			controllers.GetSuggestionWithLocationBySuggestionId)
		// Delete suggestion by suggestion id (admin)
		v1.DELETE("/suggestions/:suggestion_id",
			middleware.AuthorizeJWT([]string{"Admin"}),
			controllers.DeleteSuggestionBySuggestionId)
		// Get all suggestions by tourist username (admin, tourist themselves)
		v1.GET("/suggestions/tourist/:tourist_username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}),
			controllers.GetAllSuggestionsWithLocationByTouristUsername)
		// Delete suggestions by tourist username (admin, tourist themselves)
		v1.DELETE("/suggestions/tourist/:tourist_username",
			middleware.AuthorizeJWT([]string{"Admin", "Tourist"}),
			controllers.DeleteSuggestionsByTouristUsername)

	}

	var err error
	if os.Getenv("PORT") == "9000" {
		err = router.Run(":9000")
	} else {
		err = router.RunTLS(":443", os.Getenv("CERT_PATH"), os.Getenv("KEY_PATH"))
	}
	if err != nil {
		panic(err)
	}
}
