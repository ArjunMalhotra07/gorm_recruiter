package routes

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
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
	var employerAPIs *gin.RouterGroup = router.Group("/employer")
	{
		router.Use(middlewares.JwtVerify(seeders.JwtSecret))
		router.Use(middlewares.CheckEmployer())
		EmployerRoutes(employerAPIs, driver)
	}
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
