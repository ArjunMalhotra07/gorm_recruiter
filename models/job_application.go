package models

import (
	"time"

	"gorm.io/gorm"
)

type JobApplication struct {
	UserID    uint           `json:"-"`                                        // Foreign key for User who applied
	JobID     uint           `json:"-"`                                        // Foreign key for the Job applied to
	Status    string         `gorm:"size:255;default:'pending'" json:"status"` // E.g., pending, approved, rejected
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Adds ID, CreatedAt, UpdatedAt, DeletedAt fields
}
