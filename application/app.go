package application

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/routes"
	"gorm.io/gorm"
)

func New(driver *gorm.DB) *models.App {
	var d models.Env = models.Env{driver}
	var env *models.Env = &d
	return &models.App{Router: routes.AppRoutes(env, driver), Driver: driver}
}

func StartServer(app *models.App) error {
	server := &http.Server{Addr: ":8080", Handler: app.Router}
	err := http.ListenAndServe(server.Addr, server.Handler)
	/*
		or
		err := http.ListenAndServe(":8080", app.router)
	*/
	if err != nil {
		return err
	}
	return nil
}
