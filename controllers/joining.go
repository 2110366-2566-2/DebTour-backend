package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkJoinTourRequest(joinTourRequest models.JoinTourRequest) bool {
	if joinTourRequest.TourId == 0 || len(joinTourRequest.JoinedMembers) == 0 {
		return false
	}
	for _, member := range joinTourRequest.JoinedMembers {
		if member.FirstName == "" || member.LastName == "" || member.Age <= 0 {
			return false
		}
		if len(member.FirstName) > 50 || len(member.LastName) > 50 {
			return false

		}

		// name must contain only letters
		for _, char := range member.FirstName {
			if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
				return false
			}
		}
		for _, char := range member.LastName {
			if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
				return false
			}
		}
	}
	return true
}

// JoinTour godoc
// @Summary Join tour
// @Description Join tour
// @description Role allowed: "Tourist"
// @tags joinings
// @ID JoinTour
// @Accept json
// @Produce json
// @Param joinTourRequest body models.JoinTourRequest true "Join tour request"
// @Security ApiKeyAuth
// @Success 200 {object} models.Joining
// @Router /joinings [post]
func JoinTour(c *gin.Context) {
	var joinTourRequest models.JoinTourRequest
	if err := c.BindJSON(&joinTourRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if checkJoinTourRequest(joinTourRequest) == false {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid join tour request"})
		return
	}

	touristUsername := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))
	joinTourRequest.TouristUsername = touristUsername;

	tour, err := database.GetTourByTourId(int(joinTourRequest.TourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	if tour.MemberCount+uint(len(joinTourRequest.JoinedMembers)) > tour.MaxMemberCount {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "The joining members is exceed the max member count of the tour"})
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
		if err := database.CreateJoining(&joining, database.MainDB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
	}

	tour.MemberCount += uint(len(joinTourRequest.JoinedMembers))
	if err := database.UpdateTour(&tour, database.MainDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "join tour successfully"})
}

// GetAllJoinings godoc
// @Summary Get all joinings
// @Description Get all joinings
// @description Role allowed: "Admin"
// @Tags joinings
// @ID GetAllJoinings
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Joining
// @Router /joinings [get]
func GetAllJoinings(c *gin.Context) {
	joinings, err := database.GetALlJoinings(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": true, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": joinings})
}

// GetTouristByTourId godoc
// @summary Get a tourist by tourId
// @description Get a tourist by tourId
// @description Role allowed: "Admin" and "AgencyOwner"
// @tags joinings
// @id GetTouristByTourId
// @produce json
// @param id path int true "Tour ID"
// @Security ApiKeyAuth
// @success 200 {array} models.JoinedMembers
// @router /tours/tourists/{id} [get]
func GetTouristByTourId(c *gin.Context) {

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	id := uint(id64)

	// check if tour exists
	if _, err := database.GetTourByTourId(int(id), database.MainDB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	joinedMembers, err := database.GetJoinedMembersByTourId(id, database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(joinedMembers), "data": joinedMembers})
}