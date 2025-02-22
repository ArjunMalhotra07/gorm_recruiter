package routes

import (
	"fmt"
	"net/http"

	_ "github.com/ArjunMalhotra07/gorm_recruiter/docs"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AppRoutes(config *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(ConfigMiddleware(config))
	router.GET("/", DefaultRoute)
	fmt.Println("Hey")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	var authAPIs *gin.RouterGroup = router.Group("/")
	AuthRoutes(authAPIs)
	var employerAPIs *gin.RouterGroup = router.Group("/employer")
	EmployerRoutes(employerAPIs)
	var jobsAPIs *gin.RouterGroup = router.Group("/jobs")
	JobRoutes(jobsAPIs)
	var miscAPIs *gin.RouterGroup = router.Group("/misc")
	MiscRoutes(miscAPIs)
	return router
}
func DefaultRoute(c *gin.Context) {
	message := models.Response{Message: "Hey!"}
	c.JSON(http.StatusOK, message)
}

func ConfigMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}
