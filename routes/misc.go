package routes

import (
	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/misc"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/misc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MiscRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	miscRepo := repo.NewMiscRepo(driver)
	miscHandler := handlers.NewMiscHandler(miscRepo)
	router.GET("/getall", miscHandler.GetAllApplicants)
	router.GET("/getresumes", miscHandler.GetAllResumes)
}
