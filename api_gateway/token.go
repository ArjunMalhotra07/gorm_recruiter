package apigateway

import (
	"fmt"
	"time"

	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(uuid string, isAdmin bool) (string, error) {
	//! Create a new token object
	var token *jwt.Token = jwt.New(jwt.SigningMethodHS256)
	//! Set claims (payload)
	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid                                  // Example data
	claims["exp"] = time.Now().Add(time.Hour * 700).Unix() // Token expires in 700 hours
	claims["is_admin"] = isAdmin
	//! Generate encoded token and sign it with a secret
	tokenString, err := token.SignedString([]byte(seeders.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func VerifyToken(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//! Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
