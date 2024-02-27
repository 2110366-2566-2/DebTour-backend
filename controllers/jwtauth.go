package controllers

import (
	// "DebTour/database"
	"DebTour/database"
	"DebTour/models"
	"fmt"
	"os"
	// "strings"

	// "strings"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// jwt service
type JWTService interface {
	GenerateToken(username string, role string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	IsUser   bool   `json:"isuser"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

type RoleInput struct {
	Roles string `json:"roles"`
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (controllers *jwtServices) GenerateToken(username string, role string, isUser bool) string {
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>> username:", username)
	claims := &authCustomClaims{
		username,
		role,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    controllers.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(controllers.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (controllers *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(controllers.secretKey), nil
	})

}

func ValidateTokenHandler(c *gin.Context) {
	encodedToken := c.Param("token")
	token, err := JWTAuthService().ValidateToken(encodedToken)

	if !token.Valid {
		fmt.Println("testing")
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
	} else {
		claims := token.Claims.(jwt.MapClaims)
		if claims.Valid() != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func ValidateRoleHandler(c *gin.Context) {
	encodedToken := c.Param("token")
	token, err := JWTAuthService().ValidateToken(encodedToken)

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	user, err = database.GetUserByUsername(claims["username"].(string), database.MainDB)
	// check role
	if err != nil || user.Role != claims["role"] {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid role"})
		return
	}
	// var roles []string
	// if c.ShouldBindJSON(&roles) != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid roles type"})
	// 	return
	// }
	roles := []string{"admin", "user"}
	// Roles := strings.Split(roles.Roles, ", ")
	// for _, role := range Roles {
	// 	if role == user.Role {
	// 		c.JSON(http.StatusOK, gin.H{"success": true})
	// 		return
	// 	}
	// }
	for _, role := range roles {
		if role == user.Role {
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>", role, user.Role)
			c.JSON(http.StatusOK, gin.H{"success": true})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "mismatch role"})

}
