package routes

import (
	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/employer"
	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/employer"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmployerRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	employerRepo := repo.NewEmployerRepo(driver)
	employerHandler := handlers.NewEmployerHandler(employerRepo)
	router.Use(middlewares.JwtVerify(seeders.JwtSecret))
	router.Use(middlewares.CheckEmployer())
	router.GET("/getmyjobs", employerHandler.GetMyJobsDetail)
	router.GET("/getapplicantdata", employerHandler.GetApplicantData)
	router.POST("/postjob", employerHandler.AddJob)
}
