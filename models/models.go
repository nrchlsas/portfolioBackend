package models

import (
	"encoding/json"
	"time"
	"gorm.io/datatypes"
)

// Skill model
type Skill struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Experience model
type Experience struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Company     string         `gorm:"type:varchar(255);not null" json:"company"`
	Period      string         `gorm:"type:varchar(100);not null" json:"period"`
	Description string         `gorm:"type:text;not null" json:"description"`
	Highlights  datatypes.JSON `gorm:"type:json" json:"highlights"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
func (Experience) TableName() string {
	return "experience"
}

// Service model
type Service struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Icon           string    `gorm:"type:varchar(10);not null" json:"icon"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	Gradient       string    `gorm:"type:varchar(100);not null" json:"gradient"`
	DarkGradient   string    `gorm:"type:varchar(100);not null" json:"dark_gradient"`
	BgGradient     string    `gorm:"type:varchar(100);not null" json:"bg_gradient"`
	Color          string    `gorm:"type:varchar(50);not null" json:"color"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Project model
type Project struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Title        string         `gorm:"type:varchar(255);not null" json:"title"`
	Description  string         `gorm:"type:text;not null" json:"description"`
	Tags         datatypes.JSON `gorm:"type:json" json:"tags"`
	Gradient     string         `gorm:"type:varchar(100);not null" json:"gradient"`
	Year         int            `json:"year"`
	Emoji        string         `gorm:"type:varchar(10)" json:"emoji"`
	TagGradients datatypes.JSON `gorm:"type:json" json:"tag_gradients"`
	LinkColor    string         `gorm:"type:varchar(50)" json:"link_color"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// Contact model
type Contact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string         `gorm:"type:varchar(20);not null" json:"phone"`
	Location  string         `gorm:"type:varchar(255);not null" json:"location"`
	Social    datatypes.JSON `gorm:"type:json" json:"social"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
func (Contact) TableName() string {
	return "contact"
}

// Response wrapper for API responses
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// CreateSkillRequest request body
type CreateSkillRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateSkillRequest request body
type UpdateSkillRequest struct {
	Name string `json:"name"`
}

// CreateExperienceRequest request body
type CreateExperienceRequest struct {
	Title       string   `json:"title" binding:"required"`
	Company     string   `json:"company" binding:"required"`
	Period      string   `json:"period" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Highlights  []string `json:"highlights"`
}

// UpdateExperienceRequest request body
type UpdateExperienceRequest struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Period      string   `json:"period"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
}

// CreateServiceRequest request body
type CreateServiceRequest struct {
	Icon         string `json:"icon" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Gradient     string `json:"gradient" binding:"required"`
	DarkGradient string `json:"dark_gradient" binding:"required"`
	BgGradient   string `json:"bg_gradient" binding:"required"`
	Color        string `json:"color" binding:"required"`
}

// UpdateServiceRequest request body
type UpdateServiceRequest struct {
	Icon         string `json:"icon"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Gradient     string `json:"gradient"`
	DarkGradient string `json:"dark_gradient"`
	BgGradient   string `json:"bg_gradient"`
	Color        string `json:"color"`
}

// CreateProjectRequest request body
type CreateProjectRequest struct {
	Title        string                 `json:"title" binding:"required"`
	Description  string                 `json:"description" binding:"required"`
	Tags         []string               `json:"tags"`
	Gradient     string                 `json:"gradient" binding:"required"`
	Year         int                    `json:"year"`
	Emoji        string                 `json:"emoji"`
	TagGradients json.RawMessage        `json:"tag_gradients"`
	LinkColor    string                 `json:"link_color"`
}

// UpdateProjectRequest request body
type UpdateProjectRequest struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Tags         []string        `json:"tags"`
	Gradient     string          `json:"gradient"`
	Year         int             `json:"year"`
	Emoji        string          `json:"emoji"`
	TagGradients json.RawMessage `json:"tag_gradients"`
	LinkColor    string          `json:"link_color"`
}

// ContactRequest request body
type ContactRequest struct {
	Email    string      `json:"email" binding:"required,email"`
	Phone    string      `json:"phone" binding:"required"`
	Location string      `json:"location" binding:"required"`
	Social   interface{} `json:"social"`
}
