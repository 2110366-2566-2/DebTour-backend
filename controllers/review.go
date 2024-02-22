package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllReviews godoc
// @Summary Get all reviews
// @Description Get all reviews
// @Tags review
// @Accept json
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReviews(c *gin.Context) {
	reviews, err := database.GetAllReviews(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": reviews})
}

// GetReviewById godoc
// @Summary Get review by id
// @Description Get review by id
// @Tags review
// @Accept json
// @Produce json
// @Param reviewId path int true "Review ID"
// @Success 200 {object} models.Review
// @Router /reviews/{reviewId} [get]
func GetReviewById(c *gin.Context) {
	reviewIdString := c.Param("reviewId")
	reviewId, err := strconv.ParseInt(reviewIdString, 10, 64)

	review, err := database.GetReviewById(uint(reviewId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": review})
}

// GetReviewsByTourId godoc
// @Summary Get reviews by tour id
// @Description Get reviews by tour id
// @Tags review
// @Accept json
// @Produce json
// @Param tourId path int true "Tour ID"
// @Success 200 {object} []models.Review
// @Router /reviews/tour/{tourId} [get]
func GetReviewsByTourId(c *gin.Context) {
	tourIdString := c.Param("tourId")
	tourId, err := strconv.ParseInt(tourIdString, 10, 64)

	reviews, err := database.GetReviewsByTourId(uint(tourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": reviews})
}

// CreateReview godoc
// @Summary Create review
// @Description Create review
// @Tags review
// @Accept json
// @Produce json
// @Param tourId path int true "Tour ID"
// @Param review body models.ReviewRequest true "Review"
// @Success 200 {object} models.Review
// @Router /reviews/tour/{tourId} [post]
func CreateReview(c *gin.Context) {
	tourIdString := c.Param("tourId")
	tourId, err := strconv.ParseInt(tourIdString, 10, 64)

	var reviewRequest models.ReviewRequest
	err = c.ShouldBindJSON(&reviewRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	review := reviewRequest.ToReview(uint(tourId))
	err = database.CreateReview(review, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": review})
}

// UpdateReview godoc
// @Summary Update review
// @Description Update review
// @Tags review
// @Accept json
// @Produce json
// @Param reviewId path int true "Review ID"
// @Param review body models.ReviewRequest true "Review"
// @Success 200
// @Router /reviews/{reviewId} [put]
func UpdateReview(c *gin.Context) {
	reviewIdString := c.Param("reviewId")
	reviewId, err := strconv.ParseInt(reviewIdString, 10, 64)

	var reviewRequest models.ReviewRequest
	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	review := reviewRequest.ToReview(uint(reviewId))

	_, err = database.GetReviewById(uint(reviewId), database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.UpdateReview(review, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteReview godoc
// @Summary Delete review
// @Description Delete review
// @Tags review
// @Accept json
// @Produce json
// @Param reviewId path int true "Review ID"
// @Success 200
// @Router /reviews/{reviewId} [delete]
func DeleteReview(c *gin.Context) {
	reviewIdString := c.Param("reviewId")
	reviewId, err := strconv.ParseInt(reviewIdString, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	_, err = database.GetReviewById(uint(reviewId), database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteReview(uint(reviewId), database.MainDB)
}