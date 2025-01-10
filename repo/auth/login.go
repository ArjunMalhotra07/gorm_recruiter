package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *AuthRepo) LoginUser(email, password string) (*models.User, error) {
	var currentUser models.User
	err := r.Driver.Where("email = ? AND password_hash = ?", email, password).First(&currentUser).Error
	if err != nil {
		return nil, err
	}
	return &currentUser, nil
}
