package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *JobRepo) GetJobData(jobID string) (*models.Job, error) {
	var job models.Job
	err := r.Driver.Preload("PostedBy").First(&job, "job_id = ?", jobID).Error
	if err != nil {
		return nil, err
	}
	return &job, err
}
