package employer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
)

func AddJob(env *models.Env, w http.ResponseWriter, r *http.Request) {
	claimsInterface := r.Context().Value(constants.Claims)
	if claimsInterface == nil {
		response := models.Response{Message: "Unauthorized request", Status: http.StatusUnauthorized}
		handlers.SendResponse(w, response, http.StatusUnauthorized)
		return
	}

	claims, ok := claimsInterface.(jwt.MapClaims)
	if !ok {
		response := models.Response{Message: "Invalid claims format", Status: http.StatusUnauthorized}
		handlers.SendResponse(w, response, http.StatusUnauthorized)
		return
	}
	fmt.Println(claims[constants.UniqueID].(string))
	if isEmployer, exists := claims[constants.IsEmployer].(bool); !exists || !isEmployer {
		response := models.Response{Message: "Only employers can post jobs", Status: http.StatusUnauthorized}
		handlers.SendResponse(w, response, http.StatusUnauthorized)
		return
	}
	//! Decode the incoming JSON body into a Job struct
	var currentJob models.Job
	err := json.NewDecoder(r.Body).Decode(&currentJob)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	//! Generating new id for job
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	currentJob.Uuid = string(newUUID)
	currentJob.PostedBy = claims[constants.UniqueID].(string)
	//! Add job in table
	if err := env.Create(&currentJob).Error; err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := models.Response{Message: "Job Posted successfully!", Status: 200}
	handlers.SendResponse(w, response, http.StatusOK)
}
