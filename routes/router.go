package routes

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/config"
	"github.com/gin-gonic/gin"
)

func AppRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	// router.Use(middlewares.Logger)
	router.GET("/", DefaultRoute)
	var authAPIs *gin.RouterGroup = router.Group("/")
	{
		AuthRoutes(authAPIs, cfg.MySql.Driver)
	}
	var employerAPIs *gin.RouterGroup = router.Group("/employer")
	{
		EmployerRoutes(employerAPIs, cfg.MySql.Driver)
	}
	var jobsAPIs *gin.RouterGroup = router.Group("/jobs")
	{
		JobRoutes(jobsAPIs, cfg.MySql.Driver)
	}
	var miscAPIs *gin.RouterGroup = router.Group("/misc")
	{
		MiscRoutes(miscAPIs, cfg.MySql.Driver)
	}
	return router
}
func DefaultRoute(c *gin.Context) {
	message := models.Response{Message: "Hey!"}
	c.JSON(http.StatusOK, message)
}
