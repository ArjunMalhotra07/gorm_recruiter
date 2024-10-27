package jobs

import (
	"net/http"
	"os/exec"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func ApplyToJob(env *models.Env, w http.ResponseWriter, r *http.Request) {
	jobID := chi.URLParam(r, constants.JobID)
	userID := r.Context().Value("claims").(jwt.MapClaims)[constants.UniqueID].(string)
	//! Check if job exists
	var job models.Job
	if err := env.DB.Where("job_id = ?", jobID).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Job doesn't exist or has been deleted!", Status: http.StatusNotFound}
			handlers.SendResponse(w, response, http.StatusNotFound)
			return
		}
		response := models.Response{Message: "Error checking job existence", Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	//! Generate Application ID
	applicationID, err := exec.Command("uuidgen").Output()
	if err != nil {
		response := models.Response{Message: "Error generating Application ID", Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	var application models.JobApplication
	application.ApplicationID = string(applicationID)
	application.ApplicantID = userID
	application.JobID = jobID
	//! Add application to table
	if err := env.Create(&application).Error; err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := models.Response{Message: "Job Applied successfully!", Status: 200}
	handlers.SendResponse(w, response, http.StatusOK)
}
