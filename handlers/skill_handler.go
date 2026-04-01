package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/models"
)

// GetSkills godoc
// @Summary Get all skills
// @Description Get list of all skills
// @Tags Skills
// @Produce json
// @Success 200 {object} models.Response
// @Router /api/skills [get]
func GetSkills(c *gin.Context) {
	var skills []models.Skill
	result := config.GetDB().Find(&skills)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to fetch skills",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Skills fetched successfully",
		Data:    skills,
	})
}

// GetSkillByID godoc
// @Summary Get a skill by ID
// @Description Get a skill details by ID
// @Tags Skills
// @Produce json
// @Param id path int true "Skill ID"
// @Success 200 {object} models.Response
// @Router /api/skills/{id} [get]
func GetSkillByID(c *gin.Context) {
	id := c.Param("id")
	var skill models.Skill
	
	result := config.GetDB().First(&skill, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Skill not found",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Skill fetched successfully",
		Data:    skill,
	})
}

// CreateSkill godoc
// @Summary Create a new skill
// @Description Create a new skill
// @Tags Skills
// @Accept json
// @Produce json
// @Param request body models.CreateSkillRequest true "Skill request"
// @Success 201 {object} models.Response
// @Router /api/skills [post]
func CreateSkill(c *gin.Context) {
	var req models.CreateSkillRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	skill := models.Skill{Name: req.Name}
	result := config.GetDB().Create(&skill)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create skill",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Skill created successfully",
		Data:    skill,
	})
}

// UpdateSkill godoc
// @Summary Update a skill
// @Description Update skill details
// @Tags Skills
// @Accept json
// @Produce json
// @Param id path int true "Skill ID"
// @Param request body models.UpdateSkillRequest true "Skill request"
// @Success 200 {object} models.Response
// @Router /api/skills/{id} [put]
func UpdateSkill(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateSkillRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	var skill models.Skill
	if result := config.GetDB().First(&skill, id); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Skill not found",
			Error:   result.Error.Error(),
		})
		return
	}

	result := config.GetDB().Model(&skill).Updates(req)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update skill",
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Skill updated successfully",
		Data:    skill,
	})
}

// DeleteSkill godoc
// @Summary Delete a skill
// @Description Delete a skill by ID
// @Tags Skills
// @Produce json
// @Param id path int true "Skill ID"
// @Success 200 {object} models.Response
// @Router /api/skills/{id} [delete]
func DeleteSkill(c *gin.Context) {
	id := c.Param("id")
	
	result := config.GetDB().Delete(&models.Skill{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to delete skill",
			Error:   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Skill not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Skill deleted successfully",
	})
}
