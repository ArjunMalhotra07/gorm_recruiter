package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	// Primary key with UUID; use the gorm:"type:char(36)" to set it as a UUID with the appropriate length
	Uuid         string           `gorm:"type:char(36);primaryKey" json:"uuid"`
	Title        string           `gorm:"size:255;not null" json:"title"`        // Set a size limit and make it non-null
	Description  string           `gorm:"type:text;not null" json:"description"` // Use text type for large content
	CompanyName  string           `gorm:"size:255;not null" json:"company_name"` // Make non-null with a length constraint
	PostedBy     string           `gorm:"size:255;not null" json:"posted_by"`    // Record the poster’s name or ID
	IsActive     bool             `gorm:"default:true;column:status" json:"active_status"`
	Applications []JobApplication `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE;" json:"applications"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}
