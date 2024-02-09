package controllers

import (
	"DebTour/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

// FilterTours godoc
// @Summary Filter tours
// @Description Filter tours
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	}

	if offset == "" {
		offsetInt = 0
	} else {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	}

	fmt.Println(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo, limitInt, offsetInt)

	tours, err := models.FilterTours(name, startDate, endDate, overviewLocation, memberCountFrom, memberCountTo, priceFrom, priceTo, offsetInt, limitInt)

	type FilterToursDto struct {
		TourId           int     `json:"tourId"`
		TourName         string  `json:"tourName"`
		StartDate        string  `json:"startDate"`
		EndDate          string  `json:"endDate"`
		OverviewLocation string  `json:"overviewLocation"`
		MemberCount      uint    `json:"memberCount"`
		MaxMemberCount   uint    `json:"maxMemberCount"`
		Price            float64 `json:"price"`
	}
	var toursDto []FilterToursDto
	for _, tour := range tours {
		toursDto = append(toursDto, FilterToursDto{
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tours": toursDto})
}
