package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// "strings"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "NOT_HERE",
		ClientID:     "SECRET",
		ClientSecret: "SECRET",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "EIEI"
	Role             = ""
)

func InitializeOauthenv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	googleOauthConfig.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")
	googleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	oauthStateString = os.Getenv("OAUTH_STATE_STRING")
}

func HandleMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandleGoogleLogin(c *gin.Context) {
	Role = c.Param("role")
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> In Redirect")
	c.Redirect(http.StatusTemporaryRedirect, url)
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Redirected")
}

func HandleGoogleCallback(c *gin.Context) {
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> In Callback")
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

	var buffer map[string]interface{}
	output := make(map[string]interface{})
	err = json.Unmarshal(contents, &buffer)
	if err != nil {
		fmt.Println("Failed to decode response body:", err)
	}
	c.Params = append(c.Params, gin.Param{Key: "username", Value: buffer["id"].(string)})
	c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	token_jwt := loginController.Login(c)
	if token_jwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}
	output["username"] = buffer["id"]
	output["email"] = buffer["email"]
	output["image"] = buffer["picture"]
	output["role"] = Role
	output["password"] = "password"
	output["phone"] = "0000000000"

	var user models.User
	jsonbyte, _ := json.Marshal(output)
	json.Unmarshal(jsonbyte, &user)
	database.CreateUser(&user, database.MainDB)

	output["firstname"] = buffer["given_name"]
	output["lastname"] = buffer["family_name"]
	output["token"] = token_jwt
	c.JSON(http.StatusOK, gin.H{"success": true, "data": output})

}

func GetProfile(c *gin.Context) {
	// Retrieve username from the Gin context set by the AuthMiddleware
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}

	// You can use the username to fetch user profile information from the database or any other source
	// For demonstration purposes, we'll simply return the username
	c.JSON(http.StatusOK, gin.H{"username": username})
}
