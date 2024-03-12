package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create function for create agency
// this is my createuser function , use this format to create
// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	var username CreateUserInput
// 	c.ShouldBindJSON(&username)
// 	fmt.Println(">>>>>>>>>>>>>>>> User: ", username.Username)
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	err := database.CreateUser(&user, database.MainDB)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
// }

//create function for create agency

// CreateAgency godoc
// @Summary Create a agency
// @Description Create a agency
// @Tags agencies
// @Accept json
// @Produce json
// @Param agency body models.Agency true "Agency"
// @Success 200 {object} models.Agency
// @Router /agencies [post]
func CreateAgency(c *gin.Context) {
	var agency models.Agency
	if err := c.ShouldBindJSON(&agency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateAgency(&agency, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
}

//create function for get all agencies

// GetAllAgencies godoc
// @Summary Get all agencies
// @Description Get all agencies
// @Tags agencies
// @Produce json
// @Success 200 {array} models.Agency
// @Router /agencies [get]
func GetAllAgencies(c *gin.Context) {
	agencies, err := database.GetAllAgencies(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(agencies), "data": agencies})
}

//create function for get agency by username

// GetAgencyByUsername godoc
// @Summary Get agency by username
// @Description Get agency by username
// @Tags agencies
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Agency
// @Router /agencies/{username} [get]
func GetAgencyByUsername(c *gin.Context) {
	username := c.Param("username")
	agency, err := database.GetAgencyByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
}

//create function for update agency

// UpdateAgency godoc
// @Summary Update a agency
// @Description Update a agency
// @Tags agencies
// @Accept json
// @Produce json
// @Param agency body models.Agency true "Agency"
// @Success 200 {object} models.Agency
// @Router /agencies [put]
func UpdateAgency(c *gin.Context) {
	var agency models.Agency
	if err := c.ShouldBindJSON(&agency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.UpdateAgency(agency, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
}

//create function for delete agency

// DeleteAgency godoc
// @Summary Delete a agency
// @Description Delete a agency
// @Tags agencies
// @Accept json
// @Produce json
// @Param agency body models.Agency true "Agency"
// @Success 200 {object} models.Agency
// @Router /agencies [delete]
func DeleteAgency(c *gin.Context) {
	var agency models.Agency
	if err := c.ShouldBindJSON(&agency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.DeleteAgency(agency, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
}
