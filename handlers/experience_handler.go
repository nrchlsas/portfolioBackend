package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/models"
	"gorm.io/datatypes"
)

// GetExperience godoc
// @Summary Get all experience
// @Description Get list of all experience
// @Tags Experience
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/experience [get]
func GetExperience(c *gin.Context) {
	var experiences []models.Experience
	result := config.GetDB().Find(&experiences)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to fetch experience",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Experience fetched successfully",
		Data:    experiences,
	})
}

// GetExperienceByID godoc
// @Summary Get experience by ID
// @Description Get experience details by ID
// @Tags Experience
// @Produce json
// @Param id path int true "Experience ID"
// @Success 200 {object} models.Response
// @Router /api/experience/{id} [get]
func GetExperienceByID(c *gin.Context) {
	id := c.Param("id")
	var experience models.Experience
	
	result := config.GetDB().First(&experience, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Experience not found",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Experience fetched successfully",
		Data:    experience,
	})
}

// CreateExperience godoc
// @Summary Create a new experience
// @Description Create a new experience entry
// @Tags Experience
// @Accept json
// @Produce json
// @Param request body models.CreateExperienceRequest true "Experience request"
// @Success 201 {object} models.Response
// @Router /api/experience [post]
func CreateExperience(c *gin.Context) {
	var req models.CreateExperienceRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	highlightsJSON, err := json.Marshal(req.Highlights)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid highlights format",
			Error:   err.Error(),
		})
		return
	}

	experience := models.Experience{
		Title:       req.Title,
		Company:     req.Company,
		Period:      req.Period,
		Description: req.Description,
		Highlights:  datatypes.JSON(highlightsJSON),
	}

	result := config.GetDB().Create(&experience)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create experience",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Experience created successfully",
		Data:    experience,
	})
}

// UpdateExperience godoc
// @Summary Update an experience
// @Description Update experience details
// @Tags Experience
// @Accept json
// @Produce json
// @Param id path int true "Experience ID"
// @Param request body models.UpdateExperienceRequest true "Experience request"
// @Success 200 {object} models.Response
// @Router /api/experience/{id} [put]
func UpdateExperience(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateExperienceRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	var experience models.Experience
	if result := config.GetDB().First(&experience, id); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Experience not found",
			Error:   result.Error.Error(),
		})
		return
	}

	updateData := map[string]interface{}{}
	if req.Title != "" {
		updateData["title"] = req.Title
	}
	if req.Company != "" {
		updateData["company"] = req.Company
	}
	if req.Period != "" {
		updateData["period"] = req.Period
	}
	if req.Description != "" {
		updateData["description"] = req.Description
	}
	if len(req.Highlights) > 0 {
		highlightsJSON, _ := json.Marshal(req.Highlights)
		updateData["highlights"] = datatypes.JSON(highlightsJSON)
	}

	result := config.GetDB().Model(&experience).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update experience",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Experience updated successfully",
		Data:    experience,
	})
}

// DeleteExperience godoc
// @Summary Delete an experience
// @Description Delete an experience by ID
// @Tags Experience
// @Produce json
// @Param id path int true "Experience ID"
// @Success 200 {object} models.Response
// @Router /api/experience/{id} [delete]
func DeleteExperience(c *gin.Context) {
	id := c.Param("id")
	
	result := config.GetDB().Delete(&models.Experience{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to delete experience",
			Error:   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Experience not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Experience deleted successfully",
	})
}
