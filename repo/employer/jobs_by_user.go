package repo

import (
	"log"
	"strings"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func (r *EmployerRepo) GetJobsPostedByUser(userID string, isActive bool) ([]models.Job, error) {
	var jobs []models.Job
	userID = strings.TrimSpace(userID)
	err := r.Driver.Where("posted_by_id = ? AND is_active = ?", userID, true).Find(&jobs).Error
	log.Printf("Fetched jobs: %v, Error: %v", jobs, err)
	log.Printf("userID %v", userID)
	return jobs, err
}
