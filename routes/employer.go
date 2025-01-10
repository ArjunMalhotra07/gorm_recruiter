package routes

import (
	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/employer"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/employer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmployerRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	employerRepo := repo.NewEmployerRepo(driver)
	employerHandler := handlers.NewEmployerHandler(employerRepo)
	router.GET("/getmyjobs", employerHandler.GetMyJobsDetail)
}
