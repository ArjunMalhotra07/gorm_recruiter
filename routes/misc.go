package routes

import (
	"net/http"

	handlers "github.com/ArjunMalhotra07/gorm_recruiter/handlers/misc"
	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/config"
	repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/misc"
	"github.com/gin-gonic/gin"
)

func MiscRoutes(router *gin.RouterGroup) {
	router.GET("/getall", miscHandlerWrapper(func(miscHandler *handlers.MiscHandler, c *gin.Context) {
		miscHandler.GetAllApplicants(c)
	}))

	router.GET("/getresumes", miscHandlerWrapper(func(miscHandler *handlers.MiscHandler, c *gin.Context) {
		miscHandler.GetAllResumes(c)
	}))
}

func miscHandlerWrapper(handlerFunc func(*handlers.MiscHandler, *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, exists := c.Get("config")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
			return
		}

		config := cfg.(*config.Config)
		miscRepo := repo.NewMiscRepo(config.MySql.Driver)
		miscHandler := handlers.NewMiscHandler(miscRepo)

		handlerFunc(miscHandler, c)
	}
}
