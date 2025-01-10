package middlewares

import (
	"fmt"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckEmployer() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsInterface, exists := c.Get(constants.Claims)
		if !exists {
			response := models.Response{Message: "Unauthorized request", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		if claimsInterface == nil {
			response := models.Response{Message: "Unauthorized request", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			response := models.Response{Message: "Invalid claims format", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		if isEmployer, exists := claims[constants.IsEmployer].(bool); !exists || !isEmployer {
			response := models.Response{Message: "Not an employer!", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			return
		} else {
			fmt.Println(claims[constants.IsEmployer].(bool))
			c.Set("claims", claims)
			c.Next()
		}
	}
}
