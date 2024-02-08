package controllers

import (
	"DebTour/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTours godoc
// @summary Get all tours
// @description Get all tours
// @id GetAllTours
// @produce json
// @success 200 {array} models.Tour
// @router /tours [get]
func GetAllTours(c *gin.Context) {

	tours, err := models.GetAllTours()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tours})
}

// CreateTour godoc
// @summary Create a tour
// @description Create a tour with the input JSON data
// @id CreateTour
// @accept json
// @produce json
// @success 200 {object} models.Tour
// @router /tours [post]
func CreateTour(c *gin.Context) {

	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err := models.CreateTour(&tour)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tour})
}

// UpdateTour godoc
// @summary Update a tour
// @description Update a tour with the input JSON data
// @id UpdateTour
// @accept json
// @produce json
// @success 200 {object} models.Tour
// @router /tours [put]
func UpdateTour(c *gin.Context) {

	var tour models.Tour
	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err := models.UpdateTour(&tour)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tour})
}

// DeleteTour godoc
// @summary Delete a tour
// @description Delete a tour
// @id DeleteTour
// @produce json
// @param id path int true "Tour ID"
// @success 200 {string} string
// @router /tours/{tourId} [delete]
func DeleteTour(c *gin.Context) {

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	id := uint(id64)
	err = models.DeleteTour(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tour deleted successfully"})
}
