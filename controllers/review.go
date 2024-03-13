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
// @tags reviews
// @Produce  json
// @Success 200 {array} models.Review
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
// @tags reviews
// @Produce  json
// @Param id path int true "Review ID"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [get]
func GetReviewById(c *gin.Context) {
	reviewId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
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
// @tags reviews
// @Produce  json
// @Param id path int true "Tour ID"
// @Success 200 {array} models.Review
// @Router /reviews/tour/{id} [get]
func GetReviewsByTourId(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	reviews, err := database.GetReviewsByTourId(uint(tourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": reviews})
}

// GetAverageRatingByTourId godoc
// @Summary Get average rating by tour id
// @Description Get average rating by tour id
// @tags reviews
// @Produce  json
// @Param id path int true "Tour ID"
// @Success 200 {number} float64
// @Router /reviews/averageRating/{id} [get]
func GetAverageRatingByTourId(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	reviews, err := database.GetReviewsByTourId(uint(tourId), database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	averageRating := 0.0
	for _, review := range reviews {
		averageRating += float64(review.RatingScore)
	}
	averageRating = averageRating / float64(len(reviews))

	c.JSON(http.StatusOK, gin.H{"success": true, "data": averageRating})

}

// GetReviewsByTouristUsername godoc
// @Summary Get reviews by tourist username
// @Description Get reviews by tourist username
// @tags reviews
// @Produce  json
// @Param username path string true "Tourist Username"
// @Success 200 {array} models.Review
// @Router /reviews/tourist/{username} [get]
func GetReviewsByTouristUsername(c *gin.Context) {
	username := c.Param("username")
	reviews, err := database.GetReviewsByTouristUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": reviews})
}

// CreateReview godoc
// @Summary Create a review
// @Description Create a review
// @tags reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Tour ID"
// @Param review body models.ReviewRequest true "Review"
// @Success 200 {object} models.Review
// @Router /reviews/tour/{id} [post]
func CreateReview(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var reviewRequest models.ReviewRequest
	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	review := reviewRequest.ToReview(uint(tourId))
	err = database.CreateReview(review, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": review})
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review
// @tags reviews
// @Produce  json
// @Param id path int true "Review ID"
// @Success 200
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	reviewId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteReview(uint(reviewId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteReviewsByTourId godoc
// @Summary Delete reviews by tour id
// @Description Delete reviews by tour id
// @tags reviews
// @Produce  json
// @Param id path int true "Tour ID"
// @Success 200
// @Router /reviews/tour/{id} [delete]
func DeleteReviewsByTourId(c *gin.Context) {
	tourId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteReviewsByTourId(uint(tourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteReviewsByTouristUsername godoc
// @Summary Delete reviews by tourist username
// @Description Delete reviews by tourist username
// @tags reviews
// @Produce  json
// @Param username path string true "Tourist Username"
// @Success 200
// @Router /reviews/tourist/{username} [delete]
func DeleteReviewsByTouristUsername(c *gin.Context) {
	username := c.Param("username")
	err := database.DeleteReviewsByTouristUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}