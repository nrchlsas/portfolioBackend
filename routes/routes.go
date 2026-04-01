package routes

import (
	"github.com/gin-gonic/gin"
	"golang-backend/handlers"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// API v1 routes
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", handlers.HealthCheck)

		// Skills routes
		skills := api.Group("/skills")
		{
			skills.GET("", handlers.GetSkills)
			skills.GET("/:id", handlers.GetSkillByID)
			skills.POST("", handlers.CreateSkill)
			skills.PUT("/:id", handlers.UpdateSkill)
			skills.DELETE("/:id", handlers.DeleteSkill)
		}

		// Experience routes
		experience := api.Group("/experience")
		{
			experience.GET("", handlers.GetExperience)
			experience.GET("/:id", handlers.GetExperienceByID)
			experience.POST("", handlers.CreateExperience)
			experience.PUT("/:id", handlers.UpdateExperience)
			experience.DELETE("/:id", handlers.DeleteExperience)
		}

		// Services routes
		services := api.Group("/services")
		{
			services.GET("", handlers.GetServices)
			services.GET("/:id", handlers.GetServiceByID)
			services.POST("", handlers.CreateService)
			services.PUT("/:id", handlers.UpdateService)
			services.DELETE("/:id", handlers.DeleteService)
		}

		// Projects routes
		projects := api.Group("/projects")
		{
			projects.GET("", handlers.GetProjects)
			projects.GET("/:id", handlers.GetProjectByID)
			projects.POST("", handlers.CreateProject)
			projects.PUT("/:id", handlers.UpdateProject)
			projects.DELETE("/:id", handlers.DeleteProject)
		}

		// Contact routes
		contact := api.Group("/contact")
		{
			contact.GET("", handlers.GetContact)
			contact.PUT("", handlers.UpdateContact)
		}
	}
}
