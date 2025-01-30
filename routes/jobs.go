package routes

import (
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/jobs"
	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/config"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/jobs"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JwtVerify(seeders.JwtSecret))

	router.POST("/apply/:job_id", jobHandlerWrapper(func(jobHandler *handlers.JobsHandler, c *gin.Context) {
		jobHandler.ApplyToJob(c)
	}))

	router.GET("/", jobHandlerWrapper(func(jobHandler *handlers.JobsHandler, c *gin.Context) {
		jobHandler.GetAllJobs(c)
	}))

	router.GET("/jobdata/:job_id", jobHandlerWrapper(func(jobHandler *handlers.JobsHandler, c *gin.Context) {
		jobHandler.GetJobData(c)
	}))

	router.POST("/uploadresume", jobHandlerWrapper(func(jobHandler *handlers.JobsHandler, c *gin.Context) {
		jobHandler.UploadResume(c)
	}))
}

func jobHandlerWrapper(handlerFunc func(*handlers.JobsHandler, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, exists := c.Get("config")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
			return
		}

		config := cfg.(*config.Config)
		jobRepo := repo.NewJobRepo(config.MySql.Driver)
		jobHandler := handlers.NewJobHandler(jobRepo)

		handlerFunc(jobHandler, c)
	}
}
