package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-backend/config"
	_ "golang-backend/docs"
	"golang-backend/models"
	"golang-backend/routes"
)

// @title Portfolio Backend API
// @version 1.0
// @description Portfolio Backend RESTful API with Swagger documentation
// @termsOfService http://swagger.io/terms/

// @contactName API Support
// @contactUrl http://www.example.com/support
// @contactEmail support@example.com

// @licenseName Apache 2.0
// @licenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	// Initialize database
	_, err = config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	// Auto migrate tables
	db := config.GetDB()
	err = db.AutoMigrate(
		&models.Skill{},
		&models.Experience{},
		&models.Service{},
		&models.Project{},
		&models.Contact{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(corsMiddleware())

	// Setup Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.SetupRoutes(router)

	// Get server port from environment or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	fmt.Printf("Swagger docs available at http://localhost:%s/swagger/index.html\n", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Placeholder for Swagger spec endpoints
func init() {
	// Swagger UI will be available at /swagger/index.html
	// This is handled by swaggo/gin-swagger
}
