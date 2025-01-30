package models

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
)

type App struct {
	Router http.Handler
	Config *config.Config
}
