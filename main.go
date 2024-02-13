package main

import (
	"DebTour/controllers"
	"DebTour/models"
	"os"

	"github.com/gin-contrib/cors"
	// "go/token"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"DebTour/docs"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/zalando/gin-oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"

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
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// func SetupOauth() {
// 	controllers.InitializeOauthenv()
// }

var redirectURL, credFile string // new
func init() {
	bin := path.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s
================
`, bin)
		flag.PrintDefaults()
	}
	flag.StringVar(&redirectURL, "redirect", "http://localhost:9000/api/v1/auth/", "URL to be redirected to after authorization.")
	flag.StringVar(&credFile, "cred-file", "./test-clientid.google.json", "Credential JSON file")
}

var secret []byte
var sessionName string
var stateKey string
var sessionID string

func main() {
	flag.Parse()
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	}
	stateKey = "state"
	sessionID = "ginoauth_google_session"
	secret = []byte("secret")
	sessionName = "goquestsession"
	google.Setup(redirectURL, credFile, scopes, secret)

	models.InitDB()

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
		v1.GET("/tours/tourists/:id", controllers.GetTouristByTourId)
		v1.POST("/tours", controllers.CreateTour)
		v1.PUT("/tours/:id", controllers.UpdateTour)
		v1.DELETE("/tours/:id", controllers.DeleteTour)
		v1.GET("/tours/filter", controllers.FilterTours)

		v1.GET("/activities", controllers.GetAllActivities)

		v1.POST("/joinings", controllers.JoinTour)
		v1.GET("/joinings", controllers.GetAllJoinings)
		// v1.GET("/", controllers.HandleMain)
		// v1.GET("/GoogleLogin", controllers.HandleGoogleLogin)
		// v1.GET("/GoogleCallback", controllers.HandleGoogleCallback)
		// v1.GET("/", controllers.HandleMain)
		// v1.GET("/GoogleLogin", controllers.HandleGoogleLogin)
		// v1.GET("/GoogleCallback", controllers.HandleGoogleCallback)

		v1.GET("/login", google.LoginHandler)
		v1.Use(google.Auth())
		v1.GET("/auth", UserInfoHandler)
		v1.GET("/api", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello from private for groups"})
		})
	}

	router.Run(":9000")
}

func UserInfoHandler(ctx *gin.Context) { // new
	var (
		res goauth.Userinfo
		ok  bool
	)

	val := ctx.MustGet("user")
	if res, ok = val.(goauth.Userinfo); !ok {
		res = goauth.Userinfo{Name: "no user"}
	}

	// ctx.JSON(http.StatusOK, gin.H{"Hello": "from private", "user": res.Email})

	output := make(map[string]interface{})
	output["username"] = res.Id
	output["email"] = res.Email
	output["firstname"] = res.GivenName
	output["lastname"] = res.FamilyName
	output["picture"] = res.Picture


	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": output})
}
