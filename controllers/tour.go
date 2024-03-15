package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTours godoc
// @summary Get all tours
// @description Get all tours
// @tags tours
// @id GetAllTours
// @produce json
// @success 200 {array} models.Tour
// @router /tours [get]
func GetAllTours(c *gin.Context) {

	tours, err := database.GetAllTours(database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(tours), "data": tours})
}

// GetTourByID godoc
// @Summary Get tour by id
// @Description Get tour by id
// @Tags tours
// @ID GetTourByID
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {object} models.TourWithActivitiesWithLocation
// @Router /tours/{id} [get]
func GetTourByID(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}
	tourActivityLocation, err := database.GetTourWithActivitiesWithLocationByTourId(id, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourActivityLocation})
}

// GetTouristByTourId godoc
// @summary Get a tourist by tourId
// @description Get a tourist by tourId
// @tags tours
// @id GetTouristByTourId
// @produce json
// @param id path int true "Tour ID"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
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

// CreateTour godoc
// @summary Create a tour
// @description Create a tour with the input JSON data
// @tags tours
// @id CreateTour
// @accept json
// @produce json
// @param tour body models.TourWithActivitiesWithLocationRequest true "Tour"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @success 200 {object} models.TourWithActivitiesWithLocation
// @router /tours [post]
func CreateTour(c *gin.Context) {

	tx := database.MainDB.Begin()

	var tourWithActivitiesWithLocationRequest models.TourWithActivitiesWithLocationRequest
	if err := c.ShouldBindJSON(&tourWithActivitiesWithLocationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tour, err := models.ToTour(tourWithActivitiesWithLocationRequest, 0, "dummyAgency")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}
	err = database.CreateTour(&tour, tourWithActivitiesWithLocationRequest.Activities, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tourWithActivitiesWithLocation, err := database.GetTourWithActivitiesWithLocationByTourId(int(tour.TourId), tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourWithActivitiesWithLocation})
}

// UpdateTour godoc
// @summary Update a tour
// @description Update a tour with the input JSON data
// @tags tours
// @id UpdateTour
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param tour body models.Tour true "Tour"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @success 200 {object} string
// @router /tours/{id} [put]
func UpdateTour(c *gin.Context) {

	tourId, err := strconv.Atoi(c.Param("id"))

	// check if tour exists
	if _, err := database.GetTourByTourId(tourId, database.MainDB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tour, err := database.GetTourByTourId(tourId, database.MainDB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	tour.TourId = uint(tourId)

	err = database.UpdateTour(&tour, database.MainDB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tour updated successfully"})
}

// DeleteTour godoc
// @summary Delete a tour
// @description Delete a tour
// @tags tours
// @id DeleteTour
// @produce json
// @param id path int true "Tour ID"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @success 200 {string} string
// @router /tours/{id} [delete]
func DeleteTour(c *gin.Context) {

	tx := database.MainDB.Begin()

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	id := uint(id64)

	// check if tour exists
	if _, err := database.GetTourByTourId(int(id), tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		tx.Rollback()
		return
	}

	err = database.DeleteTour(id, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tour deleted successfully"})
}

// FilterTours godoc
// @Summary Filter tours
// @Description Filter tours
// @Tags tours
// @ID FilterTours
// @Produce json
// @Param name query string false "Name"
// @Param startDate query string false "Start date"
// @Param endDate query string false "End date"
// @Param overviewLocation query string false "Overview location"
// @Param memberCountFrom query string false "Member count from"
// @Param memberCountTo query string false "Member count to"
// @Param priceFrom query string false "Price from"
// @Param priceTo query string false "Price to"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {array} models.Tour
// @Router /tours/filter [get]
func FilterTours(c *gin.Context) {
	name := c.Query("name")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	overviewLocation := c.Query("overviewLocation")
	memberCountFrom := c.Query("memberCountFrom")
	memberCountTo := c.Query("memberCountTo")
	priceFrom := c.Query("priceFrom")
	priceTo := c.Query("priceTo")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}
	if overviewLocation == "" {
		overviewLocation = "%"
	} else {
		overviewLocation = "%" + overviewLocation + "%"
	}
	if startDate == "" {
		startDate = "1970-01-01"
	}
	if endDate == "" {
		endDate = "3000-01-01"
	}
	if memberCountFrom == "" {
		memberCountFrom = "0"
	}
	if memberCountTo == "" {
		memberCountTo = strconv.Itoa(math.MaxInt)
	}
	if priceFrom == "" {
		priceFrom = "0"
	}
	if priceTo == "" {
		priceTo = strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64)
	}

	var limitInt int
	var offsetInt int
	var err error
	if limit == "" {
		limitInt = -1
	} else {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid limit"})
			return
		}
	}

	if offset == "" {
		offsetInt = 0
	} else {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid offset"})
			return
		}
	}

	fmt.Println(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo, limitInt, offsetInt)

	tours, err := database.FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo, offsetInt, limitInt, database.MainDB)

	var filteredToursResponse []models.FilteredToursResponse
	for _, tour := range tours {
		filteredToursResponse = append(filteredToursResponse, models.FilteredToursResponse{
			TourId:           int(tour.TourId),
			TourName:         tour.Name,
			StartDate:        tour.StartDate,
			EndDate:          tour.EndDate,
			OverviewLocation: tour.OverviewLocation,
			MemberCount:      tour.MemberCount,
			MaxMemberCount:   tour.MaxMemberCount,
			Price:            tour.Price,
		})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(filteredToursResponse), "data": filteredToursResponse})
}

// UpdateTourActivities godoc
// @summary Update activities by tourId
// @description Update activities by tourId
// @tags tour-activities
// @id UpdateActivitiesByTourId
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activitiesWithLocation body []models.ActivityWithLocation true "Activities with location"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @success 200 {string} string
// @router /tours/activities/{id} [put]
func UpdateTourActivities(c *gin.Context) {

	tx := database.MainDB.Begin()

	tourId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	var activitiesWithLocation []models.ActivityWithLocation
	if err := c.ShouldBindJSON(&activitiesWithLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	err = database.UpdateActivitiesByTourId(uint(tourId), &activitiesWithLocation, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Activities updated successfully"})
}

// CreateTourActivities godoc
// @summary Create activities for a tour
// @description Create activities for a tour
// @tags tour-activities
// @id CreateTourActivities
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activitiesWithLocationRequest body []models.ActivityWithLocationRequest true "Activities with location request"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <token>"
// @success 200 {object} models.TourWithActivitiesWithLocation
// @router /tours/activities/{id} [post]
func CreateTourActivities(c *gin.Context) {
	tx := database.MainDB.Begin()

	tourId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	var activitiesWithLocationRequest []models.ActivityWithLocationRequest
	if err := c.ShouldBindJSON(&activitiesWithLocationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	err = database.CreateTourActivities(uint(tourId), activitiesWithLocationRequest, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tourWithActivitiesWithLocation, err := database.GetTourWithActivitiesWithLocationByTourId(tourId, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourWithActivitiesWithLocation})
}
