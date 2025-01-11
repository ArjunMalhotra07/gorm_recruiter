package routes

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRoutes(env *models.Env, driver *gorm.DB) *gin.Engine {
	router := gin.Default()
	// router.Use(middlewares.Logger)
	router.GET("/", DefaultRoute)
	var authAPIs *gin.RouterGroup = router.Group("/")
	{
		AuthRoutes(authAPIs, driver)
	}
	var employerAPIs *gin.RouterGroup = router.Group("/employer")
	{
		EmployerRoutes(employerAPIs, driver)
	}
	var jobsAPIs *gin.RouterGroup = router.Group("/jobs")
	{
		JobRoutes(jobsAPIs, driver)
	}
	var miscAPIs *gin.RouterGroup = router.Group("/misc")
	{
		MiscRoutes(miscAPIs, driver)
	}
	return router
}
func DefaultRoute(c *gin.Context) {
	message := models.Response{Message: "Hey!"}
	c.JSON(http.StatusOK, message)
}
