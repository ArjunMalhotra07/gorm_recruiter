package middlewares

import (
	"net/http"
	"strings"

	apigateway "github.com/ArjunMalhotra07/gorm_recruiter/api_gateway"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtVerify(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			response := models.Response{Message: "Missing Authorization header", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := strings.TrimSpace(strings.Replace(authorizationHeader, "Bearer", "", 1))
		token, err := apigateway.VerifyToken(tokenString, secret)
		if err != nil {
			response := models.Response{Message: "Invalid Token", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
			c.Next()
		} else {
			response := models.Response{Message: "Invalid Token", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
	}
}
