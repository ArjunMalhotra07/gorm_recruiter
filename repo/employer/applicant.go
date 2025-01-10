package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func (repo *EmployerRepo) FetchApplicantByID(applicantID string) (*models.User, error) {
	var applicant models.User
	if err := repo.Driver.Where("user_id = ?", applicantID).First(&applicant).Error; err != nil {
		return nil, err
	}
	return &applicant, nil
}
