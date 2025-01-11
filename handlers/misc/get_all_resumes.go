package misc

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func GetAllResumes(env *models.Env, w http.ResponseWriter, r *http.Request) {
	//! Fetch data from DB
	var resumes []models.Resume
	if err := env.DB.Preload("Educations").Preload("Experiences").Find(&resumes).Error; err != nil {
		response := models.Response{Message: "Error fetching resumes"}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := models.Response{
		Message: "resumes fetched successfully!",
		Data:    resumes,
	}
	handlers.SendResponse(w, response, http.StatusOK)
}
