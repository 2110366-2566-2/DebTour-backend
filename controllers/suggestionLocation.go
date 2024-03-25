package controllers

import (
	"DebTour/database"
	//"DebTour/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSuggestionLocationBySuggestionId(c *gin.Context) {
	_suggestionId := c.Param("suggestion_id")
	suggestionId, err := strconv.Atoi(_suggestionId)
	suggestionLocation, err := database.GetSuggestionLocationBySuggestionId(uint(suggestionId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": suggestionLocation})
}
