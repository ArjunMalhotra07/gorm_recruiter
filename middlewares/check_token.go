package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	apigateway "github.com/ArjunMalhotra07/gorm_recruiter/api_gateway"
	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtVerify(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := models.Response{Message: "Missing Authorization header", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		fmt.Println(tokenString)
		token, err := apigateway.VerifyToken(tokenString, secret)
		if err != nil || !token.Valid {
			response := models.Response{Message: "Invalid or expired Token", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims[constants.IsEmployer])
		fmt.Println(claims[constants.UniqueID])
		if !ok {
			response := models.Response{Message: "Invalid Token Claims", Status: http.StatusUnauthorized, Data: ""}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
