package routes

import (
	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/auth"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.RouterGroup, driver *gorm.DB) {
	authRepo := repo.NewAuthRepo(driver)
	authHandler := handlers.NewAuthHandler(authRepo)
	router.POST("/signup", authHandler.SignUp)
	router.POST("/login", authHandler.LogIn)
}
