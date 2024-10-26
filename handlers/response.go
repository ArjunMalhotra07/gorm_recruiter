package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
)

func SendResponse(w http.ResponseWriter, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
