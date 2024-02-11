package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld godoc
// @summary Hello World
// @description Just reply Hello World!
// @id HelloWorld
// @produce json
// @response 200 {string} string "Hello World!"
// @router /hello [get]
func HelloWorld(c *gin.Context) {
	fmt.Println("Hello World!")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!!!|!",
	})
}
