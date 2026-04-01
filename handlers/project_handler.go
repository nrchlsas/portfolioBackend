package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/models"
	"gorm.io/datatypes"
)

// GetProjects godoc
// @Summary Get all projects
// @Description Get list of all projects
// @Tags Projects
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/projects [get]
func GetProjects(c *gin.Context) {
	var projects []models.Project
	result := config.GetDB().Find(&projects)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to fetch projects",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Projects fetched successfully",
		Data:    projects,
	})
}

// GetProjectByID godoc
// @Summary Get a project by ID
// @Description Get a project details by ID
// @Tags Projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Response
// @Router /api/projects/{id} [get]
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	
	result := config.GetDB().First(&project, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Project not found",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Project fetched successfully",
		Data:    project,
	})
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Param request body models.CreateProjectRequest true "Project request"
// @Success 201 {object} models.Response
// @Router /api/projects [post]
func CreateProject(c *gin.Context) {
	var req models.CreateProjectRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid tags format",
			Error:   err.Error(),
		})
		return
	}

	project := models.Project{
		Title:        req.Title,
		Description:  req.Description,
		Tags:         datatypes.JSON(tagsJSON),
		Gradient:     req.Gradient,
		Year:         req.Year,
		Emoji:        req.Emoji,
		TagGradients: datatypes.JSON(req.TagGradients),
		LinkColor:    req.LinkColor,
	}

	result := config.GetDB().Create(&project)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create project",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Project created successfully",
		Data:    project,
	})
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update project details
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param request body models.UpdateProjectRequest true "Project request"
// @Success 200 {object} models.Response
// @Router /api/projects/{id} [put]
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateProjectRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	var project models.Project
	if result := config.GetDB().First(&project, id); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Project not found",
			Error:   result.Error.Error(),
		})
		return
	}

	updateData := map[string]interface{}{}
	if req.Title != "" {
		updateData["title"] = req.Title
	}
	if req.Description != "" {
		updateData["description"] = req.Description
	}
	if len(req.Tags) > 0 {
		tagsJSON, _ := json.Marshal(req.Tags)
		updateData["tags"] = datatypes.JSON(tagsJSON)
	}
	if req.Gradient != "" {
		updateData["gradient"] = req.Gradient
	}
	if req.Year > 0 {
		updateData["year"] = req.Year
	}
	if req.Emoji != "" {
		updateData["emoji"] = req.Emoji
	}
	if len(req.TagGradients) > 0 {
		updateData["tag_gradients"] = datatypes.JSON(req.TagGradients)
	}
	if req.LinkColor != "" {
		updateData["link_color"] = req.LinkColor
	}

	result := config.GetDB().Model(&project).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update project",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Project updated successfully",
		Data:    project,
	})
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by ID
// @Tags Projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Response
// @Router /api/projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	
	result := config.GetDB().Delete(&models.Project{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to delete project",
			Error:   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Project not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Project deleted successfully",
	})
}
