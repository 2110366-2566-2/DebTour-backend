package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTourImages godoc
// @Summary Get tour images
// @Description Get tour images
// @description Role allowed: everyone
// @Tags tour-images
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {array} models.TourImagesResponse "Tour images"
// @Router /tours/images/{id} [get]
func GetTourImages(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	images, err := database.GetTourImages(uint(tourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to get tour images"})
		return
	}

	imagesResponse := models.TourImagesResponse{}
	imagesResponse.TourId = uint(tourId)
	for _, image := range images {
		imagesResponse.Images = append(imagesResponse.Images, image)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": imagesResponse})
}

// CreateTourImagesByTourId godoc
// @Summary Create tour images
// @Description Create tour images
// @description Role allowed: "AgencyOwner"
// @Tags tour-images
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Param request body models.TourImagesRequest true "Tour images request"
// @Security ApiKeyAuth
// @Success 201 {string} string "Tour images created successfully"
// @Router /tours/images/{id} [post]
func CreateTourImagesByTourId(c *gin.Context) {
	var tourImagesRequest models.TourImagesRequest

	if err := c.ShouldBindJSON(&tourImagesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	tourId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	tx := database.MainDB.Begin()

	for _, imageb64 := range tourImagesRequest.Images {
		// Convert encoded-base64 image to binary
		image, err := base64.StdEncoding.DecodeString(imageb64)
		tourImage := models.TourImage{TourId: uint(tourId), Image: image}
		err = database.CreateTourImage(&tourImage, tx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": "Tour images created successfully"})
}

// DeleteTourImagesByTourId godoc
// @Summary Delete tour image
// @Description Delete tour images
// @description Role allowed: "Admin" and "AgencyOwner"
// @Tags tour-images
// @Produce json
// @Param id path int true "Tour ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "Tour images deleted successfully"
// @Router /tours/images/{id} [delete]
func DeleteTourImagesByTourId(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	err = database.DeleteTourImages(uint(tourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to delete tour images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tour images deleted successfully"})
}
