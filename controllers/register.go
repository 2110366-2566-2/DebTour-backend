package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register tourist -> create user then create tourist
// RegisterTourist godoc
// @Summary Register tourist
// @Description Create a new tourist and user
// @Tags register
// @Accept  json
// @Produce  json
// @Param user body User true "User object"
// @Success 200 {object} models.Tourist
func RegisterTourist(c *gin.Context) {
	tx := database.MainDB.Begin()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	var tourist models.Tourist
	if err := c.ShouldBindJSON(&tourist); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = database.CreateTourist(&tourist, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	// data := struct {
	// 	User    models.User    `json:"user"`
	// 	Tourist models.Tourist `json:"tourist"`
	// }{
	// 	User:    user,
	// 	Tourist: tourist,
	// }
	//merge data from user and tourist
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
	tx.Commit()
}

func RegisterAgency(c *gin.Context) {
	tx := database.MainDB.Begin()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err := database.CreateUser(&user, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	var agency models.Agency
	agency.Username = user.Username
	err = database.CreateAgency(&agency, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": agency})
	tx.Commit()
}
