package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllSuggestions godoc
// @Summary Get all suggestions
// @Description Get all suggestions
// @description Role allowed: "Admin"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Suggestion "List of suggestions"
// @Router /suggestions [get]
func GetAllSuggestions(c *gin.Context) {
	suggestions, err := database.GetAllSuggestions(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(suggestions), "data": suggestions})
}

// GetAllSuggestionWithLocation godoc
// @Summary Get all suggestions with location
// @Description Get all suggestions with location
// @description Role allowed: "Admin" and "Agency"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.SuggestionWithLocation "List of suggestions with location"
// @Router /suggestions/location [get]
func GetAllSuggestionsWithLocation(c *gin.Context) {
	suggestionsWithLocation, err := database.GetAllSuggestionsWithLocation(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(suggestionsWithLocation), "data": suggestionsWithLocation})
}

// GetSuggestionBySuggestionId godoc
// @Summary Get suggestion by suggestion id
// @Description Get suggestion by suggestion id
// @description Role allowed: "Admin"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Param suggestion_id path int true "Suggestion ID"
// @Success 200 {object} models.Suggestion "Suggestion"
// @Router /suggestions/{suggestion_id} [get]
func GetSuggestionBySuggestionId(c *gin.Context) {
	_suggestionId := c.Param("suggestion_id")
	suggestionId, err := strconv.Atoi(_suggestionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	suggestion, err := database.GetSuggestionBySuggestionId(uint(suggestionId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": suggestion})
}

// GetSuggestionWithLocationBySuggestionId godoc
// @Summary Get suggestion with location by suggestion id
// @Description Get suggestion with location by suggestion id
// @description Role allowed: "Admin" and "Agency"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Param suggestion_id path int true "Suggestion ID"
// @Success 200 {object} models.SuggestionWithLocation "Suggestion with location"
// @Router /suggestions/location/{suggestion_id} [get]
func GetSuggestionWithLocationBySuggestionId(c *gin.Context) {
	_suggestionId := c.Param("suggestion_id")
	suggestionId, err := strconv.Atoi(_suggestionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	suggestionWithLocation, err := database.GetSuggestionWithLocationBySuggestionId(uint(suggestionId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": suggestionWithLocation})
}

// GetAllSuggestionsWithLocationByTouristUsername godoc
// @Summary Get all suggestionsWithLocation by tourist username
// @Description Get all suggestions by tourist username
// @description Role allowed: "Admin" and "Tourist"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Param tourist_username path string true "Tourist Username"
// @Success 200 {array} models.SuggestionWithLocation "List of suggestions"
// @Router /suggestions/tourist/{tourist_username} [get]
func GetAllSuggestionsWithLocationByTouristUsername(c *gin.Context) {
	touristUsername := c.Param("tourist_username")
	suggestions, err := database.GetAllSuggestionsWithLocationByTouristUsername(touristUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(suggestions), "data": suggestions})
}

func GetSuggestionsByTouristUsername(c *gin.Context) {
	touristUsername := c.Param("tourist_username")
	suggestions, err := database.GetSuggestionsByTouristUsername(touristUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": suggestions})
}

// CreateSuggestion godoc
// @Summary Create suggestion
// @Description Create suggestion
// @description Role allowed: "Tourist"
// @Tags suggestion
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param suggestion body models.SuggestionRequest true "Suggestion"
// @Success 200 {object} models.Suggestion "Suggestion"
// @Router /suggestions [post]
func CreateSuggestion(c *gin.Context) {
	tx := database.MainDB.Begin()
	var suggestionRequest models.SuggestionRequest
	if err := c.ShouldBindJSON(&suggestionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// create suggestion
	suggestion := models.ToSuggestion(suggestionRequest)

	// check if touristUsername exists
	_, err := database.GetUserByUsername(suggestion.TouristUsername, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid touristUsername"})
		return
	}

	err = database.CreateSuggestion(suggestion, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// create location from SuggestionRequest.LocationRequest
	location := models.ToLocation(suggestionRequest.LocationRequest, 0)
	// create location
	err = database.CreateLocation(&location, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// create suggestion location
	suggestionLocation := models.SuggestionLocation{
		SuggestionId: suggestion.SuggestionId,
		LocationId:   location.LocationId,
	}
	err = database.CreateSuggestionLocation(&suggestionLocation, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// return suggestion with location
	suggestionWithLocation := models.ToSuggestionWithLocation(*suggestion, location)

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": suggestionWithLocation})
}

// DeleteSuggestionBySuggestionId godoc
// @Summary Delete suggestion by suggestion id
// @Description Delete suggestion by suggestion id
// @description Role allowed: "Admin"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Param suggestion_id path int true "Suggestion ID"
// @Success 200 {string} string "Suggestion deleted successfully"
// @Router /suggestions/{suggestion_id} [delete]
func DeleteSuggestionBySuggestionId(c *gin.Context) {
	tx := database.MainDB.Begin()
	_suggestionId := c.Param("suggestion_id")
	suggestionId, err := strconv.Atoi(_suggestionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// delete location
	suggestionLocation, err := database.GetSuggestionLocationBySuggestionId(uint(suggestionId), tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteLocation(suggestionLocation.LocationId, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// delete suggestionLocation
	err = database.DeleteSuggestionLocation(uint(suggestionId), tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// delete suggestion
	err = database.DeleteSuggestion(uint(suggestionId), tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Suggestion deleted successfully"})
}

// DeleteSuggestionsByTouristUsername godoc
// @Summary Delete suggestions by tourist username
// @Description Delete suggestions by tourist username
// @description Role allowed: "Admin" and "Tourist"
// @Tags suggestion
// @Produce json
// @Security ApiKeyAuth
// @Param tourist_username path string true "Tourist Username"
// @Success 200 {string} string "Suggestion deleted successfully"
// @Router /suggestions/tourist/{tourist_username} [delete]
func DeleteSuggestionsByTouristUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	touristUsername := c.Param("tourist_username")

	// get suggestions by tourist username
	suggestions, err := database.GetSuggestionsByTouristUsername(touristUsername, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	// delete location in suggestions
	for _, suggestion := range suggestions {
		suggestionLocation, err := database.GetSuggestionLocationBySuggestionId(suggestion.SuggestionId, tx)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		err = database.DeleteLocation(suggestionLocation.LocationId, tx)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		// delete suggestionLocation
		err = database.DeleteSuggestionLocation(suggestion.SuggestionId, tx)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
	}

	// delete suggestion
	err = database.DeleteSuggestionByTouristUsername(touristUsername, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Suggestion deleted successfully"})
}
