package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	oauthStateString = "random"
)

func InitializeOauthenv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	googleOauthConfig.RedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")
	googleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
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
