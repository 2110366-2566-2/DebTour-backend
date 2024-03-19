package middleware

import (
	"DebTour/controllers"
	"fmt"

	// "fmt"
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(roles []string, arg ...int) gin.HandlerFunc {
	usernameCheck := false
	tourOwnerCheck := false
	if len(arg) > 0 {
		usernameCheck = ((arg[0] >> 0) & 1) == 1
		tourOwnerCheck = ((arg[0] >> 1) & 1) == 1
	}

	return func(c *gin.Context) {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>> check ", usernameCheck, " ", tourOwnerCheck)
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		if _, ok := controllers.Blacklist[tokenString]; ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is logged out"})
			return
		}

		token, err := controllers.JWTAuthService().ValidateToken(tokenString)
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> username :", claims["username"].(string))
		var user models.User
		var tour models.Tour
		user, err = database.GetUserByUsername(claims["username"].(string), database.MainDB)
		// check role
		if user.Role == "sudo" {
			return
		}
		if usernameCheck && user.Role != "Admin" {
			if user.Username != c.Param("username") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "mismatch username"})
				return
			}
		}
		if tourOwnerCheck && user.Role != "Admin" {
			tourStr := c.Param("id")
			tourID, _ := strconv.Atoi(tourStr)
			tour, err = database.GetTourByTourId(tourID, database.MainDB)
			if tour.AgencyUsername != user.Username {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "mismatch tour owner"})
				return
			}
		}
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>", user.Role, " ", err.Error())
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid role"})
			return
		}
		for _, role := range roles {
			if role == user.Role {
				// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>", role, user.Role)

				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "mismatch role"})
	}
}
