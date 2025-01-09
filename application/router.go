package application

import (
	"encoding/json"
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/auth"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers/employer"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers/jobs"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers/misc"
	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/auth"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func AppRoutes(env *models.Env, driver *gorm.DB) *chi.Mux {
	var router *chi.Mux = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", DefaultRoute)
	router.Route("/", func(r chi.Router) {
		AuthRoutes(r, driver)
	})
	router.Route("/employer", func(r chi.Router) {
		r.Use(middlewares.JwtVerify(seeders.JwtSecret))
		r.Use(middlewares.CheckEmployer())
		EmployerRoutes(r, env)
	})
	router.Route("/jobs", func(r chi.Router) {
		r.Use(middlewares.JwtVerify(seeders.JwtSecret))
		JobRoutes(r, env)
	})
	router.Route("/misc", func(r chi.Router) {
		r.Use(middlewares.JwtVerify(seeders.JwtSecret))
		GeneralRoutes(r, env)
	})
	return router
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	message := models.Response{Message: "Hey!", Status: 200}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func AuthRoutes(router chi.Router, driver *gorm.DB) {
	authRepo := repo.NewAuthRepo(driver)
	authHandler := handlers.NewAuthHandler(authRepo)
	router.Post("/signup", authHandler.SignUp)
	router.Post("/login", authHandler.LogIn)
}

func EmployerRoutes(router chi.Router, env *models.Env) {
	router.Post("/postjob", func(w http.ResponseWriter, r *http.Request) {
		employer.AddJob(env, w, r)
	})
	router.Get("/getapplicantdata", func(w http.ResponseWriter, r *http.Request) {
		employer.GetApplicantData(env, w, r)
	})
	router.Get("/getmyjobsdetail", func(w http.ResponseWriter, r *http.Request) {
		employer.GetMyJobsDetail(env, w, r)
	})
}

func JobRoutes(router chi.Router, env *models.Env) {
	router.Get("/jobdata/{job_id}", func(w http.ResponseWriter, r *http.Request) {
		jobs.GetJobData(env, w, r)
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		jobs.GetAllJobs(env, w, r)
	})
	router.Post("/uploadresume", func(w http.ResponseWriter, r *http.Request) {
		jobs.UploadResume(env, w, r)
	})
	router.Post("/apply/{job_id}", func(w http.ResponseWriter, r *http.Request) {
		jobs.ApplyToJob(env, w, r)
	})
}

func GeneralRoutes(router chi.Router, env *models.Env) {
	router.Get("/getall", func(w http.ResponseWriter, r *http.Request) {
		misc.GetAllApplicants(env, w, r)
	})
	router.Get("/getresumes", func(w http.ResponseWriter, r *http.Request) {
		misc.GetAllResumes(env, w, r)
	})
}
