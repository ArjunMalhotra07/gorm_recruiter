package models

import (
	"gorm.io/gorm"
)

type User struct {
	// Primary key with UUID
	Uuid            string `gorm:"type:char(36);primaryKey" json:"uuid"`
	Name            string `gorm:"size:255;not null" json:"name"`              // Limit name to 255 chars, non-null
	Email           string `gorm:"size:255;uniqueIndex;not null" json:"email"` // Unique, non-null, and limited to 255 chars
	PasswordHash    string `gorm:"size:255;not null" json:"password_hash"`     // Password hash, non-null
	IsAdmin         bool   `gorm:"default:false" json:"is_admin"`              // Default to false for regular users
	ProfileHeadline string `gorm:"size:255" json:"profile_headline"`           // Limit headline to 255 chars
	Address         string `gorm:"type:text" json:"address"`                   // Use TEXT for potentially longer address field
	// Embedding gorm.Model for CreatedAt, UpdatedAt, DeletedAt fields
	gorm.Model
}
