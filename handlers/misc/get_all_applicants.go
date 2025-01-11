package misc

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func GetAllApplicants(env *models.Env, w http.ResponseWriter, r *http.Request) {
	//! Fetch data from DB
	var applicants []models.User
	if err := env.DB.Where("is_employer = ?", false).Find(&applicants).Error; err != nil {
		response := models.Response{Message: "Error fetching applicants"}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := models.Response{
		Message: "Applicants fetched successfully!",
		Data:    applicants,
	}
	handlers.SendResponse(w, response, http.StatusOK)
}
