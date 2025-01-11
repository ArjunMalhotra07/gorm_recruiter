package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *JobRepo) GetAllJobs() (*[]models.Job, error) {
	var jobs *[]models.Job
	err := r.Driver.Where("is_active = ?", true).Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
