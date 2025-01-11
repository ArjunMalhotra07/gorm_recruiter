package repo

import (
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/gorm"
)

type MiscRepository interface {
	GetAllApplicants() (*[]models.User, error)
	GetAllResumes() ([]*models.Resume, error)
}

type MiscRepo struct {
	Driver *gorm.DB
}

func NewMiscRepo(driver *gorm.DB) *MiscRepo {
	return &MiscRepo{Driver: driver}
}
