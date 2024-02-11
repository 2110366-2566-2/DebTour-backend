package controllers

import (
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllActivities(c *gin.Context) {
	activities, err := models.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": activities})
}
