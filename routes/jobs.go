package routes

import (
	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/jobs"
	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/jobs"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JobRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	jobRepo := repo.NewJobRepo(driver)
	jobHandler := handlers.NewJobHandler(jobRepo)
	router.Use(middlewares.JwtVerify(seeders.JwtSecret))
	router.POST("/apply/:job_id", jobHandler.ApplyToJob)
	router.GET("/", jobHandler.GetAllJobs)
	router.GET("/jobdata/:job_id", jobHandler.GetJobData)
	router.POST("/uploadresume", jobHandler.UploadResume)
}
