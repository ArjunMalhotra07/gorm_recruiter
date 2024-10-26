package models

import (
	"time"

	"gorm.io/gorm"
)

// Resume struct with relationships to Education and Experience
type Resume struct {
	ResumeID          string         `gorm:"type:char(36);primaryKey" json:"resume_id"`
	UserID            string         `gorm:"uniqueIndex" json:"-"` // Foreign key for User, unique to ensure one-to-one relationship
	Skills            string         `gorm:"type:text" json:"skills"`
	ResumeFileAddress string         `gorm:"size:255" json:"resume_file_address"`
	Name              string         `gorm:"size:255;not null" json:"name"`
	Email             string         `gorm:"size:255;not null" json:"email"`
	Phone             string         `gorm:"size:15" json:"phone"`
	Educations        []Education    `gorm:"foreignKey:ResumeID" json:"educations"`  // One-to-many relationship
	Experiences       []Experience   `gorm:"foreignKey:ResumeID" json:"experiences"` // One-to-many relationship
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
}

// Education struct associated with a Resume
type Education struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	ResumeID string `json:"-"` // Foreign key for Resume
}

// Experience struct associated with a Resume
type Experience struct {
	gorm.Model
	Title        string `gorm:"size:255" json:"title"`
	Organization string `gorm:"size:255" json:"organization"`
	ResumeID     string `json:"-"` // Foreign key for Resume
}
