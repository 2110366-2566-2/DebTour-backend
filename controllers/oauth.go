package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Blacklist = make(map[string]bool)

// HandleGoogleLogout godoc
// @Summary Handle Logout
// @Description Handle Logout
// @description Role allowed: "Admin", "Agency" and "Tourist"
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /auth/logout [get]
func HandleGoogleLogout(c *gin.Context) {
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

// Login godoc
// @Summary First contact
// @Description First contact of user when login to the system
// @description Role allowed: everyone
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

func GetUsernameByTokenWithBearer(tokenWithBearer string) string {
	const BEARER_SCHEMA = "Bearer "
	token, err := JWTAuthService().ValidateToken(tokenWithBearer[len(BEARER_SCHEMA):])
	if !token.Valid {
		return err.Error()
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["username"].(string)
}
