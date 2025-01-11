package repo

import (
	"mime/multipart"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/gorm"
)

type JobsRepository interface {
	CheckIfJobExists(jobID string) error
	CheckIfApplied(applicantID, jobID string) (*models.JobApplication, error)
	CreateApplication(applicationID, applicantID, jobID string) error
	GetAllJobs() (*[]models.Job, error)
	GetJobData(jobID string) (*models.Job, error)
	SaveResumeToDirectory(userResume multipart.File, header *multipart.FileHeader) (string, error)
	ReadFileContent(resumeFilePath string) ([]byte, error)
	ParseResume(fileContent []byte) (*models.ResumeResponse, error)
	SaveParsedResumeToDatabase(parsedResume *models.Resume) error
}

type JobRepo struct {
	Driver *gorm.DB
}

func NewJobRepo(driver *gorm.DB) *JobRepo {
	return &JobRepo{Driver: driver}
}
