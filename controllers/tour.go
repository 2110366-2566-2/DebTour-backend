package controllers

import (
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

	tours, err := models.GetAllTours()

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
// @Success 200 {object} models.TourActivityLocation
// @Router /tours/{id} [get]
func GetTourByID(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}
	tourActivityLocation, err := models.GetTourById(id)
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
// @success 200 {array} models.JoinedMembers
// @router /tours/tourists/{id} [get]
func GetTouristByTourId(c *gin.Context) {

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	id := uint(id64)
	joinedMembers, err := models.GetJoinedMembersByTourId(id)

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
// @param tour body models.TourRequest true "Tour"
// @success 200 {object} models.Tour
// @router /tours [post]
func CreateTour(c *gin.Context) {

	var TourRequest models.TourRequest
	if err := c.ShouldBindJSON(&TourRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	tour := models.ToTour(TourRequest)
	err := models.CreateTour(&tour, TourRequest.Activities)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tour})
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
// @success 200 {object} models.Tour
// @router /tours/{id} [put]
func UpdateTour(c *gin.Context) {

	tourId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tour, err := models.GetOnlyTourById(tourId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&tour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = models.UpdateTour(&tour)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": tour})
}

// DeleteTour godoc
// @summary Delete a tour
// @description Delete a tour
// @tags tours
// @id DeleteTour
// @produce json
// @param id path int true "Tour ID"
// @success 200 {string} string
// @router /tours/{id} [delete]
func DeleteTour(c *gin.Context) {

	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	id := uint(id64)
	err = models.DeleteTour(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

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

	tours, err := models.FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo, offsetInt, limitInt)

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

	c.JSON(http.StatusOK, gin.H{"success": false, "count": len(filteredToursResponse), "data": filteredToursResponse})
}

// UpdateActivitiesByTourId godoc
// @summary Update activities by tourId
// @description Update activities by tourId
// @tags tours
// @id UpdateActivitiesByTourId
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activitiesUpdate body []models.ActivityResponse true "Activities Update"
// @success 200 {string} string
// @router /tours/activities/{id} [put]
func UpdateTourActivities(c *gin.Context) {

	tourId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	var activitiesResponse []models.ActivityResponse
	if err := c.ShouldBindJSON(&activitiesResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = models.UpdateActivitiesByTourId(uint(tourId), &activitiesResponse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": activitiesResponse})
}

// CreateTourActivities godoc
// @summary Create activities for a tour
// @description Create activities for a tour
// @tags tours
// @id CreateTourActivities
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activities body []models.ActivityResponse true "Activities"
// @success 200 {string} string
// @router /tours/activities/{id} [post]
func CreateTourActivities(c *gin.Context) {

	tourId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	var activitiesResponse []models.ActivityResponse
	if err := c.ShouldBindJSON(&activitiesResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	err = models.CreateTourActivities(uint(tourId), activitiesResponse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": activitiesResponse})
}
