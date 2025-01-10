package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *EmployerRepo) AddJob(job *models.Job) error {
	return r.Driver.Create(&job).Error
}
