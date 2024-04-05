package controllers

import (
	"DebTour/database"
	"DebTour/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"net/http"
	"os"
)

// GetAllTransactionPayments godoc
// @Summary Get all transaction payments
// @Description Get all transaction payments
// @description Role allowed: "Admin"
// @tags transactionPayments
// @ID GetAllTransactionPayments
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.FullTransactionPayment
// @Router /transactionPayments [get]
func GetAllTransactionPayments(c *gin.Context) {
	fullTransactionPayment, err := database.GetAllTransactionPayments(database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": fullTransactionPayment})
}

// GetTransactionPaymentByTourId godoc
// @Summary Get transaction payments by tourId
// @Description Get transaction payments by tourId
// @description Role allowed: "Admin" and "Agency Owner"
// @tags transactionPayments
// @ID GetTransactionPaymentByTourId
// @Produce json
// @Security ApiKeyAuth
// @Param tourId path string true "Tour ID"
// @Success 200 {array} models.FullTransactionPayment
// @Router /transactionPayments/tours/{tourId} [get]
func GetTransactionPaymentByTourId(c *gin.Context) {
	tourId := c.Param("tourId")
	fullTransactionPayment, err := database.GetTransactionPaymentByTourId(tourId, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": fullTransactionPayment})
}

// GetTransactionPaymentByTouristUsername godoc
// @Summary Get transaction payments by tourist username
// @Description Get transaction payments by tourist username
// @description Role allowed: "Admin" and "Tourist Owner"
// @tags transactionPayments
// @ID GetTransactionPaymentByTouristUsername
// @Produce json
// @Security ApiKeyAuth
// @Param username path string true "Tourist Username"
// @Success 200 {array} models.FullTransactionPayment
// @Router /transactionPayments/tourists/{username} [get]
func GetTransactionPaymentByTouristUsername(c *gin.Context) {
	username := c.Param("username")
	fullTransactionPayment, err := database.GetTransactionPaymentByTouristUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": fullTransactionPayment})
}

// GetTransactionPaymentByTransactionId godoc
// @Summary Get transaction payments by transactionId
// @Description Get transaction payments by transactionId
// @description Role allowed: "Admin" and "Tourist Owner" and "Agency Owner"
// @tags transactionPayments
// @ID GetTransactionPaymentByTransactionId
// @Produce json
// @Security ApiKeyAuth
// @Param transactionId path string true "Transaction ID"
// @Success 200 {array} models.FullTransactionPayment
// @Router /transactionPayments/{transactionId} [get]
func GetTransactionPaymentByTransactionId(c *gin.Context) {
	transactionId := c.Param("transactionId")
	fullTransactionPayments, err := database.GetTransactionPaymentByTransactionId(transactionId, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	username := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))

	user, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	if user.Role != "Admin" && user.Role != "sudo" {
		// Tourist Owner?
		if user.Role == "Tourist" {
			if user.Username != fullTransactionPayments.TouristUsername {
				// Unauthorized
				c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
				return
			}
		} else if user.Role == "Agency" {
			tour, err := database.GetTourByTourId(int(fullTransactionPayments.TourId), database.MainDB)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}
			if tour.AgencyUsername != user.Username {
				// Unauthorized
				c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": fullTransactionPayments})
}

// GetStripePublicKey godoc
// @Summary Get stripe public key
// @Description Get stripe public key
// @tags transactionPayments
// @ID GetStripePublicKey
// @Produce json
// @Success 200 {string} string "stripePublicKey"
// @Router /stripePublicKey [get]
func GetStripePublicKey(c *gin.Context) {
	stripePublicKey := os.Getenv("STRIPE_PUBLISHABLE_KEY")
	c.JSON(http.StatusOK, gin.H{"success": true, "data": stripePublicKey})
}

// StartTransactionPayment godoc
// @Summary Start transaction payment
// @Description Start transaction payment
// @description Role allowed: "Tourist"
// @tags transactionPayments
// @ID StartTransactionPayment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param transactionPaymentCreateForm body models.TransactionPaymentCreateForm true "Transaction Payment Create Form"
// @Success 200 {string} string "clientSecret"
// @Router /transactionPayments [post]
func StartTransactionPayment(c *gin.Context) {
	var transactionPaymentCreateForm models.TransactionPaymentCreateForm
	if err := c.BindJSON(&transactionPaymentCreateForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// check if tour exists
	_, err := database.GetTourByTourId(int(transactionPaymentCreateForm.TourId), database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	transactionPaymentCreateForm.Status = "Pending"

	username := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))
	transactionPaymentCreateForm.TouristUsername = username

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(transactionPaymentCreateForm.Amount * 100)),
		Currency: stripe.String(string(stripe.CurrencyTHB)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Create transaction payment
	tx := database.MainDB.Begin()
	err = database.CreateTransactionPayment(transactionPaymentCreateForm, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": pi.ClientSecret})
}

// UpdateTransactionStatus godoc
// @Summary Update transaction status
// @Description Update transaction status to "Failed" or "Success"
// @description Role allowed: "Tourist" and "Admin"
// @tags transactionPayments
// @ID ConfirmTransactionPayment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param transactionId path string true "Transaction ID"
// @Param status body string true "Status"
// @Success 200 {string} string "transactionId"
// @Router /transactionPayments/{transactionId} [put]
func UpdateTransactionStatus(c *gin.Context) {
	transactionId := c.Param("transactionId")
	var status string
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Check if status is valid
	if status != "Failed" && status != "Success" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid status"})
		return
	}

	// Check if transaction exists
	_, err := database.GetTransactionPaymentByTransactionId(transactionId, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Check if user is authorized
	username := GetUsernameByTokenWithBearer(c.GetHeader("Authorization"))
	user, err := database.GetUserByUsername(username, database.MainDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Unauthorized"})
		return
	}
	if user.Role != "Admin" && user.Role != "sudo" {
		if user.Role == "Tourist" {
			// Tourist Owner?
			var transactionPayment models.TransactionPayment
			_, err := database.GetTransactionPaymentByTransactionId(transactionId, database.MainDB)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}
			if user.Username != transactionPayment.TouristUsername {
				// Unauthorized
				c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Unauthorized"})
				return
			}
		}
	}

	tx := database.MainDB.Begin()
	if err := database.UpdateTransactionStatus(transactionId, status, tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": `"` + transactionId + `" updated to ` + status})
}

// DeleteTransactionPayment godoc
// @Summary Delete transaction payment
// @Description Delete transaction payment
// @description Role allowed: "Admin"
// @tags transactionPayments
// @ID DeleteTransactionPayment
// @Produce json
// @Security ApiKeyAuth
// @Param transactionId path string true "Transaction ID"
// @Success 200 {string} string "transactionId"
// @Router /transactionPayments/{transactionId} [delete]
func DeleteTransactionPayment(c *gin.Context) {
	transactionId := c.Param("transactionId")
	if err := database.DeleteTransactionPayment(transactionId, database.MainDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": transactionId})
}

// DeleteTransactionPaymentByTourId godoc
// @Summary Delete transaction payment by tourId
// @Description Delete transaction payment by tourId
// @description Role allowed: "Admin"
// @tags transactionPayments
// @ID DeleteTransactionPaymentByTourId
// @Produce json
// @Security ApiKeyAuth
// @Param tourId path string true "Tour ID"
// @Success 200 {string} string "tourId"
// @Router /transactionPayments/tours/{tourId} [delete]
func DeleteTransactionPaymentByTourId(c *gin.Context) {
	tourId := c.Param("tourId")
	tx := database.MainDB.Begin()
	if err := database.DeleteTransactionPaymentByTourId(tourId, tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": tourId})
}

// DeleteTransactionPaymentByTouristUsername godoc
// @Summary Delete transaction payment by tourist username
// @Description Delete transaction payment by tourist username
// @description Role allowed: "Admin"
// @tags transactionPayments
// @ID DeleteTransactionPaymentByTouristUsername
// @Produce json
// @Security ApiKeyAuth
// @Param username path string true "Tourist Username"
// @Success 200 {string} string "username"
// @Router /transactionPayments/tourists/{username} [delete]
func DeleteTransactionPaymentByTouristUsername(c *gin.Context) {
	username := c.Param("username")
	tx := database.MainDB.Begin()
	if err := database.DeleteTransactionPaymentByTouristUsername(username, tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": username})
}
