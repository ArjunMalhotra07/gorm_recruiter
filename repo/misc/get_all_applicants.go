package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *MiscRepo) GetAllApplicants() (*[]models.User, error) {
	var applicants []models.User
	err := r.Driver.Where("is_employer = ?", false).Find(&applicants).Error
	if err != nil {
		return nil, err
	}
	return &applicants, nil
}
