package repo

import (
	"os/exec"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func CreateUserID() (string, error) {
	userID, err := exec.Command("uuidgen").Output()
	return string(userID), err
}

func CreateEncryptedPassword(userPassword, passwordHash string) (string, error) {
	hashedPassword, err := Encrypt(userPassword, passwordHash)
	return hashedPassword, err
}

func (r *AuthRepo) CreateUser(user *models.User) error {
	return r.Driver.Create(&user).Error
}

func CreateJwtToken(userID string, isEmployer bool) (string, error) {
	jwtToken, err := CreateToken(userID, isEmployer)
	return jwtToken, err
}
