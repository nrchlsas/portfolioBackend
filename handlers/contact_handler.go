package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/models"
	"gorm.io/datatypes"
)

// GetContact godoc
// @Summary Get contact information
// @Description Get contact information
// @Tags Contact
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/contact [get]
func GetContact(c *gin.Context) {
	var contact models.Contact
	result := config.GetDB().First(&contact)
	
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Contact information not found",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Contact information fetched successfully",
		Data:    contact,
	})
}

// UpdateContact godoc
// @Summary Update contact information
// @Description Update contact information
// @Tags Contact
// @Accept json
// @Produce json
// @Param request body models.ContactRequest true "Contact request"
// @Success 200 {object} models.Response
// @Router /api/contact [put]
func UpdateContact(c *gin.Context) {
	var req models.ContactRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	var contact models.Contact
	result := config.GetDB().First(&contact)
	
	socialJSON, err := json.Marshal(req.Social)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid social format",
			Error:   err.Error(),
		})
		return
	}

	if result.Error != nil {
		// Create new contact if doesn't exist
		contact = models.Contact{
			Email:    req.Email,
			Phone:    req.Phone,
			Location: req.Location,
			Social:   datatypes.JSON(socialJSON),
		}
		createResult := config.GetDB().Create(&contact)
		if createResult.Error != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Failed to create contact",
				Error:   createResult.Error.Error(),
			})
			return
		}
	} else {
		// Update existing contact
		updateData := map[string]interface{}{
			"email":    req.Email,
			"phone":    req.Phone,
			"location": req.Location,
			"social":   datatypes.JSON(socialJSON),
		}
		updateResult := config.GetDB().Model(&contact).Updates(updateData)
		if updateResult.Error != nil {
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Failed to update contact",
				Error:   updateResult.Error.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Contact information updated successfully",
		Data:    contact,
	})
}

// HealthCheck godoc
// @Summary Health check
// @Description Health check endpoint
// @Tags Health
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "API is running",
		Data: map[string]string{
			"status": "healthy",
		},
	})
}
