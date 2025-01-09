package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *AuthRepo) LoginUser(email, password string) error {
	var currentUser models.User
	return r.Driver.Where("email = ? AND password_hash = ?", email, password).First(&currentUser).Error
}
