package models

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/config"
	"gorm.io/gorm"
)

type App struct {
	Router http.Handler
	Config *config.Config
}

type Env struct {
	*gorm.DB
}
