package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *EmployerRepo) GetApplicantsForJob(jobID string) ([]models.User, error) {
	var applicants []models.User
	err := r.Driver.Joins("JOIN job_applications ON job_applications.applicant_id = users.user_id").
		Where("job_applications.job_id = ?", jobID).
		Find(&applicants).Error
	return applicants, err
}
