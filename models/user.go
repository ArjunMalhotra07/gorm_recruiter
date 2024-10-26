package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Uuid            string           `gorm:"type:char(36);primaryKey" json:"uuid"`
	Name            string           `gorm:"size:255;not null" json:"name"`
	Email           string           `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash    string           `gorm:"size:255;not null" json:"password_hash"`
	IsEmployer      bool             `gorm:"default:false" json:"is_employer"` // Changed to IsEmployer
	ProfileHeadline string           `gorm:"size:255" json:"profile_headline"`
	Address         string           `gorm:"type:text" json:"address"`
	Profile         Profile          `gorm:"foreignKey:UserID"`
	Jobs            []Job            `gorm:"foreignKey:PostedByUserID"`
	Applications    []JobApplication `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}
