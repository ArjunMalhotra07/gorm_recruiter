package routes

import (
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/auth"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/signup", authHandlerWrapper(func(authHandler *handlers.AuthHandler, c *gin.Context) {
		authHandler.SignUp(c)
	}))
	router.POST("/login", authHandlerWrapper(func(authHandler *handlers.AuthHandler, c *gin.Context) {
		authHandler.LogIn(c)
	}))
}
func authHandlerWrapper(handlerFunc func(*handlers.AuthHandler, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, exists := c.Get("config")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
			return
		}
		config := cfg.(*config.Config)
		authRepo := repo.NewAuthRepo(config.MySql.Driver, config.Microservices.EmailService)
		authHandler := handlers.NewAuthHandler(authRepo)
		handlerFunc(authHandler, c)
	}
}
