package middleware

import (
	"DebTour/controllers"
	"fmt"
	// "fmt"
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := controllers.JWTAuthService().ValidateToken(tokenString)
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		var user models.User
		user, err = database.GetUserByUsername(claims["username"].(string), database.MainDB)
		// check role
		if err != nil || user.Role != claims["role"] {
			fmt.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>", user.Role, " ", claims["role"])
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
