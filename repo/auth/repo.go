package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/gorm"
)

type Authentication interface {
	CreateUserID() (string, error)
	CreateEncryptedPassword(userPassword, passwordHash string) (string, error)
	CreateUser(user *models.User) error
	CreateJwtToken(userID string, isEmployer bool) (string, error)
	LoginUser(email, password string) error
}

type AuthRepo struct {
	Driver *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{Driver: db}
}
