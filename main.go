package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func main() {
	router := gin.Default()

	router.GET("/", getHello)

	router.Run("localhost:9000")
}