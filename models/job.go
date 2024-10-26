package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	// Primary key with UUID; use the gorm:"type:char(36)" to set it as a UUID with the appropriate length
	Uuid              string           `gorm:"type:char(36);primaryKey" json:"uuid"`
	Title             string           `gorm:"size:255;not null" json:"title"`        // Set a size limit and make it non-null
	Description       string           `gorm:"type:text;not null" json:"description"` // Use text type for large content
	PostedOn          time.Time        `gorm:"autoCreateTime" json:"posted_on"`       // Automatically set time of posting
	TotalApplications int              `gorm:"default:0" json:"total_applications"`   // Default to 0 applications
	CompanyName       string           `gorm:"size:255;not null" json:"company_name"` // Make non-null with a length constraint
	PostedBy          string           `gorm:"size:255;not null" json:"posted_by"`    // Record the poster’s name or ID
	PostedByUserID    uint             `json:"-"`                                     // Foreign key for User who posted the job
	Status            bool             `gorm:"default:true" json:"active_status"`     // Default to true for active
	Applications      []JobApplication `gorm:"foreignKey:JobID"`                      // Many-to-many relationship with users via JobApplication
	CreatedAt         time.Time        `json:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at"`
	DeletedAt         gorm.DeletedAt   `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
}
