package controllers

import (
	"DebTour/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JoinTour godoc
// @Summary Join tour
// @Description Join tour
// @ID JoinTour
// @Accept json
// @Produce json
// @Success 200 {object} models.Joining
// @Router /joinings/join-tour [post]
func JoinTour(c *gin.Context) {
	type JoinTourRequest struct {
		TourId          uint   `json:"tourId"`
		TouristUsername string `json:"touristUsername"`
		JoinedMembers   []struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Age       uint   `json:"age"`
		} `json:"joinedMembers"`
	}
	var joinTourRequest JoinTourRequest
	if err := c.BindJSON(&joinTourRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		if err := models.CreateJoining(joining); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

}

// GetAllJoinings godoc
// @Summary Get all joinings
// @Description Get all joinings
// @ID GetAllJoinings
// @Produce json
// @Success 200 {array} models.Joining
// @Router /joinings [get]
func GetAllJoinings(c *gin.Context) {
	joinings, err := models.GetALlJoinings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"joinings": joinings})
}
