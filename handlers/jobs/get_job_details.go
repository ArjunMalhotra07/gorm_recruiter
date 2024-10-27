package jobs

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func GetJobData(env *models.Env, w http.ResponseWriter, r *http.Request) {
	//! Get Job ID
	jobID := chi.URLParam(r, constants.JobID)
	//! Fetch data from DB
	var job models.Job
	if err := env.DB.Where("job_id = ?", jobID).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Job doesn't exist or is either deleted!", Status: http.StatusNotFound}
			handlers.SendResponse(w, response, http.StatusNotFound)
			return
		}
		response := models.Response{Message: "Error fetching job details", Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := models.Response{
		Message: "Job fetched successfully!",
		Status:  http.StatusOK,
		Data:    job,
	}
	handlers.SendResponse(w, response, http.StatusOK)
}
