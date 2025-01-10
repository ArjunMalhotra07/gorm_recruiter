package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *EmployerRepo) GetJobsPostedByUser(userID string, isActive bool) ([]models.Job, error) {
	var jobs []models.Job
	err := r.Driver.Where("posted_by_id = ? AND is_active = ?", userID, isActive).Find(&jobs).Error
	return jobs, err
}
