package employer

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
)

func GetMyJobsDetail(env *models.Env, w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(constants.Claims).(jwt.MapClaims)[constants.UniqueID].(string)
	var jobs []models.Job
	//! Fetch jobs posted by the user that are active
	if err := env.DB.Where("posted_by_id = ? AND is_active = ?", userID, true).Find(&jobs).Error; err != nil {
		response := models.Response{Message: "Error fetching job details!", Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	//! If no jobs found, jobs will be an empty slice
	response := models.Response{
		Message: "Jobs fetched successfully!",
		Status:  http.StatusOK,
		Data:    jobs,
	}
	handlers.SendResponse(w, response, http.StatusOK)
}
