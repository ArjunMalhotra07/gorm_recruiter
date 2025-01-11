package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func (r *JobRepo) CheckIfJobExists(jobID string) error {
	var job models.Job
	err := r.Driver.Where("job_id = ?", jobID).First(&job).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *JobRepo) CheckIfApplied(applicantID, jobID string) (*models.JobApplication, error) {
	var application models.JobApplication
	err := r.Driver.Where("applicant_id = ? AND job_id = ?", applicantID, jobID).First(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *JobRepo) CreateApplication(applicationID, applicantID, jobID string) error {
	var application models.JobApplication
	application.ApplicationID = applicationID
	application.ApplicantID = applicantID
	application.JobID = jobID
	return r.Driver.Create(&application).Error
}
