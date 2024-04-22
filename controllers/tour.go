package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTours godoc
// @summary Get all tours
// @description Get all tours
// @description Role allowed: everyone
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

	type TourResponse struct {
		models.Tour
		FirstTourImage string
	}

	var toursResponse []TourResponse
	for _, tour := range tours {
		tourImage, err := database.GetTourImages(tour.TourId, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		if len(tourImage) == 0 {
			tourImage = append(tourImage, "No image")
		}
		toursResponse = append(toursResponse, TourResponse{tour, tourImage[0]})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(tours), "data": toursResponse})
}

// GetTourByID godoc
// @Summary Get tour by id
// @Description Get tour by id
// @description Role allowed: everyone
// @Tags tours
// @ID GetTourByID
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {object} models.TourWithActivitiesWithLocationWithImages
// @Router /tours/{id} [get]
func GetTourByID(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}
	tourActivityLocation, err := database.GetTourWithActivitiesWithLocationWithImagesByTourId(id, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	agency, err := database.GetAgencyByUsername(tourActivityLocation.AgencyUsername, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	type response struct {
		models.TourWithActivitiesWithLocationWithImages
		AgencyName string
	}
	responseData := response{
		TourWithActivitiesWithLocationWithImages: tourActivityLocation,
		AgencyName:                               agency.AgencyName,
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": responseData})
}

// CreateTour godoc
// @summary Create a tour
// @description Create a tour with the input JSON data
// @description Role allowed: "Agency"
// @tags tours
// @id CreateTour
// @accept json
// @produce json
// @param tour body models.TourWithActivitiesWithLocationWithImagesRequest true "Tour"
// @success 200 {object} models.TourWithActivitiesWithLocationWithImages
// @Security ApiKeyAuth
// @router /tours [post]
func CreateTour(c *gin.Context) {

	tx := database.MainDB.Begin()

	var tourWithActivitiesWithLocationWithImagesRequest models.TourWithActivitiesWithLocationWithImagesRequest
	if err := c.ShouldBindJSON(&tourWithActivitiesWithLocationWithImagesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	agencyUsername := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))

	tour, err := models.ToTour(tourWithActivitiesWithLocationWithImagesRequest, 0, agencyUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}
	err = database.CreateTour(&tour, tourWithActivitiesWithLocationWithImagesRequest.Activities, tourWithActivitiesWithLocationWithImagesRequest.Images, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tourWithActivitiesWithLocation, err := database.GetTourWithActivitiesWithLocationWithImagesByTourId(int(tour.TourId), tx)

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
// @description Role allowed: "Admin" and "AgencyOwner"
// @tags tours
// @id UpdateTour
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param tour body models.TourWithImagesRequest true "Tour"
// @Security ApiKeyAuth
// @success 200 {object} string
// @router /tours/{id} [put]
func UpdateTour(c *gin.Context) {
	tx := database.MainDB.Begin()
	tourId, err := strconv.Atoi(c.Param("id"))

	// check if tour exists
	tour, err := database.GetTourByTourId(tourId, tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid tour id"})
		return
	}

	// Bind the input JSON data to the tourWithImagesRequest struct
	var tourWithImagesRequest models.TourWithImagesRequest
	if err := c.ShouldBindJSON(&tourWithImagesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	// Update the tour with the input JSON data
	tour, err = models.ToTourFromTourWithImagesRequest(tourWithImagesRequest, uint(tourId), tour.AgencyUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	err = database.UpdateTour(&tour, tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	// Update Images
	err = database.UpdateTourImagesByTourId(uint(tourId), tourWithImagesRequest.Images, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": "Tour updated successfully"})
}

// DeleteTour godoc
// @summary Delete a tour
// @description Delete a tour
// @description Role allowed: "Admin" and "AgencyOwner"
// @tags tours
// @id DeleteTour
// @produce json
// @param id path int true "Tour ID"
// @Security ApiKeyAuth
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
// @Description Filter tours allow everyone
// @Tags tours
// @ID FilterTours
// @Produce json
// @Param name query string false "Name"
// @Param agencyUsername query string false "Agency username"
// @Param startDate query string false "Start date"
// @Param endDate query string false "End date"
// @Param overviewLocation query string false "Overview location"
// @Param memberCountFrom query string false "Member count from"
// @Param memberCountTo query string false "Member count to"
// @Param maxMemberCountFrom query string false "Max member count from"
// @Param maxMemberCountTo query string false "Max member count to"
// @Param availableMemberCountFrom query string false "Available member count from"
// @Param availableMemberCountTo query string false "Available member count to"
// @Param priceFrom query string false "Price from"
// @Param priceTo query string false "Price to"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {array} models.Tour
// @Router /tours/filter [get]
func FilterTours(c *gin.Context) {
	name := c.Query("name")
	agencyUsername := c.Query("agencyUsername")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	overviewLocation := c.Query("overviewLocation")
	memberCountFrom := c.Query("memberCountFrom")
	memberCountTo := c.Query("memberCountTo")
	availableMemberCountFrom := c.Query("availableMemberCountFrom")
	availableMemberCountTo := c.Query("availableMemberCountTo")
	maxMemberCountFrom := c.Query("maxMemberCountFrom")
	maxMemberCountTo := c.Query("maxMemberCountTo")
	priceFrom := c.Query("priceFrom")
	priceTo := c.Query("priceTo")
	limit := c.Query("limit")
	offset := c.Query("offset")

	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}
	if agencyUsername == "" {
		agencyUsername = "%"
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
	if availableMemberCountFrom == "" {
		availableMemberCountFrom = "0"
	}
	if availableMemberCountTo == "" {
		availableMemberCountTo = strconv.Itoa(math.MaxInt)
	}
	if maxMemberCountFrom == "" {
		maxMemberCountFrom = "0"
	}
	if maxMemberCountTo == "" {
		maxMemberCountTo = strconv.Itoa(math.MaxInt)
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
	//fmt.Println(maxMemberCountFrom, maxMemberCountTo)

	tours, err := database.FilterTours(name, agencyUsername, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, maxMemberCountFrom, maxMemberCountTo, availableMemberCountFrom, availableMemberCountTo, priceFrom, priceTo, offsetInt, limitInt, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	type FilteredTourWithImageResponse struct {
		models.FilteredToursResponse
		FirstTourImage string
	}

	var filteredToursResponse []FilteredTourWithImageResponse
	for _, tour := range tours {
		filteredTour := models.FilteredToursResponse{
			TourId:           int(tour.TourId),
			TourName:         tour.Name,
			StartDate:        tour.StartDate,
			EndDate:          tour.EndDate,
			OverviewLocation: tour.OverviewLocation,
			MemberCount:      tour.MemberCount,
			MaxMemberCount:   tour.MaxMemberCount,
			Price:            tour.Price,
		}

		images, err := database.GetTourImages(tour.TourId, database.MainDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		image := "No image"
		if len(images) > 0 {
			image = images[0]
		}

		filteredToursResponse = append(filteredToursResponse, FilteredTourWithImageResponse{filteredTour, image})
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
// @description Role allowed: "Admin" and "AgencyOwner"
// @tags tour-activities
// @id UpdateActivitiesByTourId
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activitiesWithLocation body []models.ActivityWithLocation true "Activities with location"
// @Security ApiKeyAuth
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
// @description Role allowed: "AgencyOwner"
// @tags tour-activities
// @id CreateTourActivities
// @accept json
// @produce json
// @param id path int true "Tour ID"
// @param activitiesWithLocationRequest body []models.ActivityWithLocationRequest true "Activities with location request"
// @success 200 {object} models.TourWithActivitiesWithLocationWithImages
// @Security ApiKeyAuth
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

	tourWithActivitiesWithLocationWithImages, err := database.GetTourWithActivitiesWithLocationWithImagesByTourId(tourId, tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		tx.Rollback()
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourWithActivitiesWithLocationWithImages})
}

// GetToursByAgencyUsername godoc
// @Summary Get tours by agency username
// @Description Get tours by agency username
// @description Role allowed: "AgencyOwner"
// @Tags tours
// @ID GetToursByAgencyUsername
// @Produce json
// @Param username path string true "Username"
// @Security ApiKeyAuth
// @Success 200 {array} models.Tour
// @Router /tours/agency/{username} [get]
func GetToursByAgencyUsername(c *gin.Context) {
	username := c.Param("username")
	tours, err := database.GetToursByAgencyUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "count": len(tours), "data": tours})
}
