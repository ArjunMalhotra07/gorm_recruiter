package application

import (
	"net/http"

	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	driver *gorm.DB
}

type Env struct {
	driver *gorm.DB
}

func New(driver *gorm.DB) *App {
	var d Env = Env{driver: driver}
	var env *Env = &d
	return &App{router: AppRoutes(env), driver: driver}
}
