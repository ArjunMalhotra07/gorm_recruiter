package models

import (
	"time"

	"gorm.io/gorm"
)

// Profile struct with relationships to Education and Experience
type Profile struct {
	Uuid              string         `gorm:"type:char(36);primaryKey" json:"uuid"`
	UserID            uint           `gorm:"uniqueIndex" json:"-"` // Foreign key for User, unique to ensure one-to-one relationship
	Skills            string         `gorm:"type:text" json:"skills"`
	ResumeFileAddress string         `gorm:"size:255" json:"resume_file_address"`
	Name              string         `gorm:"size:255;not null" json:"name"`
	Email             string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Phone             string         `gorm:"size:15" json:"phone"`
	Educations        []Education    `gorm:"foreignKey:ProfileID" json:"educations"`  // One-to-many relationship
	Experiences       []Experience   `gorm:"foreignKey:ProfileID" json:"experiences"` // One-to-many relationship
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
}

// Education struct associated with a Profile
type Education struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name      string         `gorm:"size:255" json:"name"`
	ProfileID uint           `json:"-"` // Foreign key for Profile
}

// Experience struct associated with a Profile
type Experience struct {
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	Title        string         `gorm:"size:255" json:"title"`
	Organization string         `gorm:"size:255" json:"organization"`
	ProfileID    uint           `json:"-"` // Foreign key for Profile
}
