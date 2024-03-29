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
// @description Role allowed: "Admin"
// @tags reviews
// @Produce  json
// @Security ApiKeyAuth
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
// @description Role allowed: everyone
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
// @description Role allowed: everyone
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
// @description Role allowed: everyone
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

	if len(reviews) == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "No reviews found"})
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
// @description Role allowed: "Admin" and "TouristThemselves"
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
// @description Role allowed: "Tourist"
// @tags reviews
// @Accept  json
// @Produce  json
// @Param id path int true "Tour ID"
// @Param review body models.ReviewRequest true "Review"
// @Security ApiKeyAuth
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

	touristUsername := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))
	reviewRequest.TouristUsername = touristUsername

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
// @description Role allowed: "Admin"
// @tags reviews
// @Produce  json
// @Param id path int true "Review ID"
// @Security ApiKeyAuth
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
// @description Role allowed: "Admin" and "AgencyOwner"
// @tags reviews
// @Produce  json
// @Param id path int true "Tour ID"
// @Security ApiKeyAuth
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
// @description Role allowed: "Admin" and "TouristThemselves"
// @tags reviews
// @Produce  json
// @Param username path string true "Tourist Username"
// @Security ApiKeyAuth
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
