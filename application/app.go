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
	*gorm.DB
}

func New(driver *gorm.DB) *App {
	var d Env = Env{driver}
	var env *Env = &d
	return &App{router: AppRoutes(env), driver: driver}
}

func (app *App) StartServer() error {
	server := &http.Server{Addr: ":8080", Handler: app.router}
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
