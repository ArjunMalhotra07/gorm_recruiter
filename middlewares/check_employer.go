package middlewares

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckEmployer() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsInterface, exists := c.Get(constants.Claims)
		if !exists || claimsInterface == nil {
			response := models.Response{Message: "Unauthorized request", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			response := models.Response{Message: "Invalid claims format", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		isEmployer, ok := claims[constants.IsEmployer].(bool)
		// fmt.Println(isEmployer)
		// fmt.Println("Checking if the user is an employer")
		if !ok || !isEmployer {
			response := models.Response{Message: "Not an employer!", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Next()
	}
}
