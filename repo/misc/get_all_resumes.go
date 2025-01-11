package repo

import "github.com/ArjunMalhotra07/gorm_recruiter/models"

func (r *MiscRepo) GetAllResumes() (*[]models.Resume, error) {
	var resumes []models.Resume
	err := r.Driver.Preload("Educations").Preload("Experiences").Find(&resumes).Error
	if err != nil {
		return nil, err
	}
	return &resumes, nil
}
