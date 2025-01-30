package routes

import (
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/employer"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
	"github.com/ArjunMalhotra07/gorm_recruiter/middlewares"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/employer"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
)

func EmployerRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JwtVerify(seeders.JwtSecret))
	router.Use(middlewares.CheckEmployer())
	router.GET("/getmyjobs", employerHandlerWrapper(func(employerHandler *handlers.EmployerHandler, c *gin.Context) {
		employerHandler.GetMyJobsDetail(c)
	}))
	router.GET("/getapplicantdata", employerHandlerWrapper(func(employerHandler *handlers.EmployerHandler, c *gin.Context) {
		employerHandler.GetApplicantData(c)
	}))
	router.POST("/postjob", employerHandlerWrapper(func(employerHandler *handlers.EmployerHandler, c *gin.Context) {
		employerHandler.AddJob(c)
	}))
}
func employerHandlerWrapper(handlerFunc func(*handlers.EmployerHandler, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, exists := c.Get("config")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
			return
		}

		config := cfg.(*config.Config)
		employerRepo := repo.NewEmployerRepo(config.MySql.Driver)
		employerHandler := handlers.NewEmployerHandler(employerRepo)

		handlerFunc(employerHandler, c)
	}
}
