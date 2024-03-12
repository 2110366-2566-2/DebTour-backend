package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type tttt struct {
// 	Username string `json:"username"`
// }

func TestRedir(c *gin.Context) {
	c.Params = append(c.Params, gin.Param{Key: "username", Value: "test"})
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
	// c.Redirect(http.StatusTemporaryRedirect, "www.googlew.com")

	// redirect to google
	// c.Redirect(http.StatusTemporaryRedirect, "/api/v1/testdir2")
}

func TestDir(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func TestLogin(c *gin.Context) {
	username := c.Query("username")
	email := c.Query("email")
	c.JSON(http.StatusOK, gin.H{"username": username, "email": email})
}
