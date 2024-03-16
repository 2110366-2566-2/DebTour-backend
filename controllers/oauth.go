package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	// "strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Blacklist = make(map[string]bool)

type BaseUrlInput struct {
	Url string `json:"baseurl"`
}

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
	Role             = "None"
	baseurl          BaseUrlInput
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
	baseurl.Url = "testlogin"
}

func HandleMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandleGoogleLogin(c *gin.Context) {
	// Role = c.Param("role")

	c.ShouldBindJSON(&baseurl)
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>", baseurl.Url)
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
	// c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	token_jwt := loginController.Login(c) // crap..., must move
	if token_jwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>", token_jwt)
	_, err = database.GetUserByUsername(buffer["id"].(string), database.MainDB)
	if err != nil && err.Error() == "record not found" {
		// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Redirecting to ", baseurl.Url+"/register?username="+buffer["id"].(string)+"&email="+buffer["email"].(string))
		c.Redirect(http.StatusTemporaryRedirect, baseurl.Url+"/register?username="+buffer["id"].(string)+"&email="+buffer["email"].(string))
	}
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Redirecting to ", baseurl.Url+"/login?username="+buffer["id"].(string)+"&email="+buffer["email"].(string))
	c.Redirect(http.StatusTemporaryRedirect, baseurl.Url+"/login?username="+buffer["id"].(string)+"&email="+buffer["email"].(string))

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>", isSuccess)

	// output["username"] = buffer["id"]
	// output["email"] = buffer["email"]
	// output["image"] = buffer["picture"]
	// output["role"] = Role
	// output["password"] = "password"
	// output["phone"] = "0000000000"

	// var user models.User
	// jsonbyte, _ := json.Marshal(output)
	// json.Unmarshal(jsonbyte, &user)
	// database.CreateUser(&user, database.MainDB)

	// output["firstname"] = buffer["given_name"]
	// output["lastname"] = buffer["family_name"]
	// output["token"] = token_jwt
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

//create godoc for HandleGoogleLogout

// HandleGoogleLogout godoc
// @Summary Handle Logout
// @Description Handle Logout
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /auth/logout [get]
func HandleGoogleLogout(c *gin.Context) {
	err := CheckToken(c)
	if err != nil {
		if err.Error() == "Authorization header is missing" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": false, "error": "Authorization header is missing"})
			return
		}
		if err.Error() == "Invalid authorization format" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": false, "error": "Invalid authorization format"})
			return
		}
		if err.Error() == "User is logged out" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "User is logged out"})
			return
		}
		if err.Error() == "Invalid token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid token"})
			return
		}
	}
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	// Add the token to the blacklist
	Blacklist[tokenString] = true

	// You may want to add additional logic such as removing the token from client storage
	// and performing any cleanup tasks

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// @Summary First contact
// @Description First contact of user when login to the system
// @Tags auth
// @Accept json
// @Produce json
// @Param firstContact body models.FirstContactModel true "First Contact"
// @Success 200 {object} models.FirstContactModel
// @Router /auth/login [post]
func Login(c *gin.Context) {
	// receive FirstContactModel

	var firstContact models.FirstContactModel
	err := c.ShouldBindJSON(&firstContact)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// check if user exists

	user, err := database.GetUserByUsername(firstContact.Id, database.MainDB)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "Failed to retrieve user from database"})
		return
	}

	// Generate JWT Token
	c.Params = append(c.Params, gin.Param{Key: "username", Value: user.Username})
	// c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token_jwt := loginController.Login(c) // crap..., must move
	if token_jwt == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve username from context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "token": token_jwt, "id": user.Role})

}

func GetToken(c *gin.Context) {
	username := c.Param("username")
	c.Params = append(c.Params, gin.Param{Key: "username", Value: username})
	// c.Params = append(c.Params, gin.Param{Key: "role", Value: Role})

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> Params: ", c.Params)
	var loginService LoginService = StaticLoginService()
	var jwtService JWTService = JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)
	token := loginController.Login(c)
	c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
}

func GetUsername(c *gin.Context) {
	tokenS := c.Param("token")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> tokenS :", tokenS)
	token, err := JWTAuthService().ValidateToken(tokenS)
	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> username :", claims["username"].(string))
	c.JSON(http.StatusOK, gin.H{"success": true, "username": claims["username"].(string)})
}

func CheckToken(c *gin.Context) error {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return errors.New("Authorization header is missing")
	}

	if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return errors.New("Invalid authorization format")
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	if _, ok := Blacklist[tokenString]; ok {
		return errors.New("User is logged out")
	}

	token, err := JWTAuthService().ValidateToken(tokenString)
	if err != nil || !token.Valid {
		return errors.New("Invalid token")
	}

	// Optionally, you can set the token in the context for later use
	return nil
}
