package application

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func AppRoutes(env *Env) *chi.Mux {
	var router *chi.Mux = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", DefaultRoute)
	router.Route("/", func(r chi.Router) {
		AuthRoutes(r, env)
	})
	return router
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	message := models.Response{Message: "Hey!", Status: 200}
	handlers.SendResponse(w, message)
}

func AuthRoutes(router chi.Router, env *Env) {
}
