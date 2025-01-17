package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/google/uuid"
)

func (r *AuthRepo) CreateUserID() string {
	userID := uuid.New().String()
	return userID
}

func (r *AuthRepo) CreateEncryptedPassword(userPassword, passwordHash string) (string, error) {
	hashedPassword, err := Encrypt(userPassword, passwordHash)
	return hashedPassword, err
}

func (r *AuthRepo) CreateUser(user *models.User) error {
	return r.Driver.Create(&user).Error
}

func (r *AuthRepo) CreateJwtToken(userID string, isEmployer bool) (string, error) {
	jwtToken, err := CreateToken(userID, isEmployer)
	return jwtToken, err
}
