package admin

import (
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
)

func AddJob(env *models.Env, w http.ResponseWriter, r *http.Request) {
	//! Check if request is coming from employers
	var claims jwt.MapClaims = r.Context().Value("claims").(jwt.MapClaims)
	if !claims["is_Employer"].(bool) {
		response := models.Response{Message: "Only employers can post jobs", Status: http.StatusUnauthorized}
		handlers.SendResponse(w, response, http.StatusUnauthorized)
		return
	}
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		handlers.SendResponse(w, response, http.StatusInternalServerError)
		return
	}
}
