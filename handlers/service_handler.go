package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/models"
)

// GetServices godoc
// @Summary Get all services
// @Description Get list of all services
// @Tags Services
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/services [get]
func GetServices(c *gin.Context) {
	var services []models.Service
	result := config.GetDB().Find(&services)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to fetch services",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Services fetched successfully",
		Data:    services,
	})
}

// GetServiceByID godoc
// @Summary Get a service by ID
// @Description Get a service details by ID
// @Tags Services
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} models.Response
// @Router /api/services/{id} [get]
func GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	
	result := config.GetDB().First(&service, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Service not found",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Service fetched successfully",
		Data:    service,
	})
}

// CreateService godoc
// @Summary Create a new service
// @Description Create a new service
// @Tags Services
// @Accept json
// @Produce json
// @Param request body models.CreateServiceRequest true "Service request"
// @Success 201 {object} models.Response
// @Router /api/services [post]
func CreateService(c *gin.Context) {
	var req models.CreateServiceRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	service := models.Service{
		Icon:         req.Icon,
		Title:        req.Title,
		Description:  req.Description,
		Gradient:     req.Gradient,
		DarkGradient: req.DarkGradient,
		BgGradient:   req.BgGradient,
		Color:        req.Color,
	}

	result := config.GetDB().Create(&service)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create service",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Service created successfully",
		Data:    service,
	})
}

// UpdateService godoc
// @Summary Update a service
// @Description Update service details
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Param request body models.UpdateServiceRequest true "Service request"
// @Success 200 {object} models.Response
// @Router /api/services/{id} [put]
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateServiceRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	var service models.Service
	if result := config.GetDB().First(&service, id); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Service not found",
			Error:   result.Error.Error(),
		})
		return
	}

	result := config.GetDB().Model(&service).Updates(req)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update service",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Service updated successfully",
		Data:    service,
	})
}

// DeleteService godoc
// @Summary Delete a service
// @Description Delete a service by ID
// @Tags Services
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} models.Response
// @Router /api/services/{id} [delete]
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	
	result := config.GetDB().Delete(&models.Service{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to delete service",
			Error:   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Service not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Service deleted successfully",
	})
}
