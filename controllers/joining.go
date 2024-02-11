package controllers

import (
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JoinTour godoc
// @Summary Join tour
// @Description Join tour
// @tags joinings
// @ID JoinTour
// @Accept json
// @Produce json
// @Param joinTourRequest body models.JoinTourRequest true "Join tour request"
// @Success 200 {object} models.Joining
// @Router /joinings [post]
func JoinTour(c *gin.Context) {
	var joinTourRequest models.JoinTourRequest
	if err := c.BindJSON(&joinTourRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	for _, member := range joinTourRequest.JoinedMembers {
		joining := models.Joining{
			TourId:          joinTourRequest.TourId,
			TouristUsername: joinTourRequest.TouristUsername,
			MemberFirstName: member.FirstName,
			MemberLastName:  member.LastName,
			MemberAge:       member.Age,
		}
		if err := models.CreateJoining(&joining); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "join tour successfully"})
}

// GetAllJoinings godoc
// @Summary Get all joinings
// @Description Get all joinings
// @Tags joinings
// @ID GetAllJoinings
// @Produce json
// @Success 200 {array} models.Joining
// @Router /joinings [get]
func GetAllJoinings(c *gin.Context) {
	joinings, err := models.GetALlJoinings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": true, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": joinings})
}
