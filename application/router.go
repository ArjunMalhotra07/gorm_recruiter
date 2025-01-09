package application

import (
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/auth"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRoutes(env *models.Env, driver *gorm.DB) *gin.Engine {
	router := gin.Default()
	// router.Use(middleware.Logger)
	router.GET("/", DefaultRoute)
	var authAPIs *gin.RouterGroup = router.Group("/")
	{
		AuthRoutes(authAPIs, driver)
	}
	// router.Route("/employer", func(r chi.Router) {
	// 	r.Use(middlewares.JwtVerify(seeders.JwtSecret))
	// 	r.Use(middlewares.CheckEmployer())
	// 	EmployerRoutes(r, env)
	// })
	// router.Route("/jobs", func(r chi.Router) {
	// 	r.Use(middlewares.JwtVerify(seeders.JwtSecret))
	// 	JobRoutes(r, env)
	// })
	// router.Route("/misc", func(r chi.Router) {
	// 	r.Use(middlewares.JwtVerify(seeders.JwtSecret))
	// 	GeneralRoutes(r, env)
	// })
	return router
}

func DefaultRoute(c *gin.Context) {
	message := models.Response{Message: "Hey!", Status: 200}
	c.JSON(http.StatusOK, message)
}
func AuthRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	authRepo := repo.NewAuthRepo(driver)
	authHandler := handlers.NewAuthHandler(authRepo)
	router.POST("/signup", authHandler.SignUp)
	router.POST("/login", authHandler.LogIn)
}

// func EmployerRoutes(router chi.Router, env *models.Env) {
// 	router.Post("/postjob", func(w http.ResponseWriter, r *http.Request) {
// 		employer.AddJob(env, w, r)
// 	})
// 	router.Get("/getapplicantdata", func(w http.ResponseWriter, r *http.Request) {
// 		employer.GetApplicantData(env, w, r)
// 	})
// 	router.Get("/getmyjobsdetail", func(w http.ResponseWriter, r *http.Request) {
// 		employer.GetMyJobsDetail(env, w, r)
// 	})
// }

// func JobRoutes(router chi.Router, env *models.Env) {
// 	router.Get("/jobdata/{job_id}", func(w http.ResponseWriter, r *http.Request) {
// 		jobs.GetJobData(env, w, r)
// 	})
// 	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		jobs.GetAllJobs(env, w, r)
// 	})
// 	router.Post("/uploadresume", func(w http.ResponseWriter, r *http.Request) {
// 		jobs.UploadResume(env, w, r)
// 	})
// 	router.Post("/apply/{job_id}", func(w http.ResponseWriter, r *http.Request) {
// 		jobs.ApplyToJob(env, w, r)
// 	})
// }

// func GeneralRoutes(router chi.Router, env *models.Env) {
// 	router.Get("/getall", func(w http.ResponseWriter, r *http.Request) {
// 		misc.GetAllApplicants(env, w, r)
// 	})
// 	router.Get("/getresumes", func(w http.ResponseWriter, r *http.Request) {
// 		misc.GetAllResumes(env, w, r)
// 	})
// }
