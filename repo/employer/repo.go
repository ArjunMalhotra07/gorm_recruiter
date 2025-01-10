package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/gorm"
)

type EmployerRepository interface {
	GetJobsPostedByUser(userID string, isActive bool) ([]models.Job, error)
	GetApplicantsForJob(jobID string) ([]models.User, error)
	GetApplicantData(userID string) (*models.User, error)
}

type EmployerRepo struct {
	Driver *gorm.DB
}

func NewEmployerRepo(db *gorm.DB) *EmployerRepo {
	return &EmployerRepo{Driver: db}
}
