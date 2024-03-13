package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create function for create tourist

// CreateTourist godoc
// @Summary Create a tourist
// @Description Create a tourist
// @Tags tourists
// @Accept json
// @Produce json
// @Param tourist body models.Tourist true "Tourist"
// @Success 200 {object} models.Tourist
// @Router /tourists [post]
func CreateTourist(c *gin.Context) {
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateTourist(&tourist, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}

//create function for get all tourists

// GetAllTourists godoc
// @Summary Get all tourists
// @Description Get all tourists
// @Tags tourists
// @Produce json
// @Success 200 {array} models.Tourist
// @Router /tourists [get]
func GetAllTourists(c *gin.Context) {
	tourists, err := database.GetAllTourists(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(tourists), "data": tourists})
}

//create function for get tourist by username

// GetTouristByUsername godoc
// @Summary Get tourist by username
// @Description Get tourist by username
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Tourist
// @Router /tourists/{username} [get]
func GetTouristByUsername(c *gin.Context) {
	username := c.Param("username")
	tourist, err := database.GetTouristByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourist})
}

//create function for delete tourist

// DeleteTouristByUsername godoc
// @Summary Delete tourist and user
// @Description Delete tourist and user by username
// @Tags tourists
// @Produce json
// @Param username path string true "Username"
// @Success 200 {string} string	"Tourist deleted successfully"
// @Router /tourists/{username} [delete]
func DeleteTouristByUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")

	//check is username exist
	_, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = database.DeleteUserByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.DeleteTouristByUsername(username, database.MainDB)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tourist deleted successfully"})
	tx.Commit()
}

// UpdateTouristByUsername godoc
// @Summary Update a tourist
// @Description Update a tourist and user also
// @Tags tourists
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param tourist body models.TouristWithUser true "Tourist"
// @Success 200 {object} models.TouristWithUser
// @Router /tourists/{username} [put]
func UpdateTouristByUsername(c *gin.Context) {
	tx := database.MainDB.Begin()
	username := c.Param("username")
	var payload models.TouristWithUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	var user models.User
	user.Username = payload.Username
	user.Phone = payload.Phone
	user.Email = payload.Email
	user.Image = payload.Image
	user.Role = "Tourist"

	var tourist models.Tourist
	tourist.Username = payload.Username
	tourist.CitizenId = payload.CitizenId
	tourist.FirstName = payload.FirstName
	tourist.LastName = payload.LastName
	tourist.Address = payload.Address
	tourist.BirthDate = payload.BirthDate
	tourist.Gender = payload.Gender
	tourist.DefaultPayment = payload.DefaultPayment

	err := database.UpdateUserByUsername(username, user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = database.UpdateTouristByUsername(username, tourist, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	data := gin.H{
		"username":       user.Username,
		"phone":          user.Phone,
		"email":          user.Email,
		"image":          user.Image,
		"role":           user.Role,
		"created_time":   user.CreatedTime,
		"citizenId":      tourist.CitizenId,
		"firstName":      tourist.FirstName,
		"lastName":       tourist.LastName,
		"address":        tourist.Address,
		"birthDate":      tourist.BirthDate,
		"gender":         tourist.Gender,
		"defaultPayment": tourist.DefaultPayment,
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	tx.Commit()
}
