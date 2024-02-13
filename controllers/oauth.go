package controllers

import (
	"flag"
	"fmt"

	// "io/ioutil"
	"net/http"
	"os"
	"path"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/zalando/gin-oauth2/google"

	//"encoding/json"
	"github.com/gin-gonic/gin"
	//"github.com/zalando/gin-oauth2/google"

	goauth "google.golang.org/api/oauth2/v2"
)

var redirectURL, credFile string // new
var secret []byte
var sessionName string
var stateKey string
var sessionID string
var scopes []string

func InitOauth() string {
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
	scopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	}
	flag.Parse()
	stateKey = "state"
	sessionID = "ginoauth_google_session"
	secret = []byte("secret")
	sessionName = "goquestsession"
	google.Setup(redirectURL, credFile, scopes, secret)
	return sessionName
}

// UserInfoHandler godoc
// @Summary Get user info
// @Description Get user info
// @Tags oauth
// @ID UserInfoHandler
// @Produce  json
// @Success 200 {object} string
// @Router /oauth/userinfo [get]
func UserInfoHandler(c *gin.Context) {

	var (
		userInfo goauth.Userinfo
		ok       bool
	)

	val := c.MustGet("user")
	if userInfo, ok = val.(goauth.Userinfo); !ok {
		userInfo = goauth.Userinfo{Name: "no user"}
	}

	//get error message
	// err := c.MustGet("error")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
	// 	return
	// }

	output := make(map[string]interface{})
	output["username"] = userInfo.Id
	output["email"] = userInfo.Email
	output["firstname"] = userInfo.GivenName
	output["lastname"] = userInfo.FamilyName
	output["picture"] = userInfo.Picture

	c.JSON(http.StatusOK, gin.H{"success": true, "data": output})
}

// LogoutHandler godoc
// @Summary Logout
// @Description Logout
// @Tags oauth
// @ID LogoutHandler
// @Produce  json
// @Success 200 {string} string
// @Router /oauth/logout [get]
func LogoutHandler(c *gin.Context) {
	// Clear the user's session data
	ClearSessionData(c)

	// Optionally, revoke the access token
	//revokeAccessToken(c)

	// Redirect the user to the login page
	c.Redirect(http.StatusFound, "/api/v1/login")
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "You have been logged out"})
}

// ClearSessionData godoc
// @Summary Clear session data
// @Description Clear session data
// @Tags oauth
// @ID ClearSessionData
// @Produce  json
// @Success 200 {string} string
// @Router /oauth/clear-session [get]
func ClearSessionData(c *gin.Context) {
	// You need to implement session clearing based on your session management mechanism
	// This could involve clearing cookies, session tokens, or any other session data

	// For example, if you are using sessions with Gin's session middleware:
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// func GetTokenHandler(c *gin.Context) {
// 	// Get the OAuth 2.0 token from the request header or query parameter, depending on how it's passed
// 	token := c.GetHeader("Authorization") // Example: Get token from Authorization header

// 	// Return the token in the response
// 	c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
// }

// GetUserData godoc
// @Summary Get user data
// @Description Get user data
// @Tags oauth
// @ID GetUserData
// @Produce  json
// @Success 200 {object} string
// @Router /oauth/getuserdata [get]
func GetUserData(c *gin.Context) {
	val := c.MustGet("user")
	userInfo, ok := val.(goauth.Userinfo)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "unable to retrieve user"})
		return
	}

	// Construct the response containing user information
	output := make(map[string]interface{})
	output["username"] = userInfo.Id
	output["email"] = userInfo.Email
	output["firstname"] = userInfo.GivenName
	output["lastname"] = userInfo.FamilyName
	output["picture"] = userInfo.Picture
	c.JSON(http.StatusOK, gin.H{"success": true, "data": output})
}

// GetUserSession godoc
// @Summary Get user session
// @Description Get user session
// @Tags oauth
// @ID GetUserSession
// @Produce  json
// @Success 200 {string} string
// @Router /oauth/getusersession [get]
func GetUserSession(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	//get error message
	err := session.Get("error")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

// SetUserSession godoc
// @Summary Set user session
// @Description Set user session
// @Tags oauth
// @ID SetUserSession
// @Produce  json
// @Success 200 {string} string
// @Router /oauth/setusersession [get]
func SetUserSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", "DejeD")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Session data set successfully"})
}
