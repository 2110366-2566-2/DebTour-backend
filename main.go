package main

import (
	"DebTour/controllers"
	"DebTour/models"

	"github.com/gin-gonic/gin"

	"DebTour/docs"
	"io/ioutil"
	"net/http"

	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	//"github.com/gorilla/pat"
	//"github.com/gorilla/sessions"
	"fmt"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/GoogleCallback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLINET_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random"
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
	router.LoadHTMLFiles("index.html")
	SetUpSwagger()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", HandleMain)
	router.GET("/GoogleLogin", HandleGoogleLogin)
	router.GET("/GoogleCallback", HandleGoogleCallback)
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
	}

	router.Run(":9000")
}

func HandleMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("Failed to fetch user info:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.String(http.StatusOK, "Content: %s\n", contents)
}
