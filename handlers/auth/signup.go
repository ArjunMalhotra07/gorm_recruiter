package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	apigateway "github.com/ArjunMalhotra07/gorm_recruiter/api_gateway"
	"github.com/ArjunMalhotra07/gorm_recruiter/handlers"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// ! This should be in an env file in production
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// ! Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	fmt.Println(secretKey)
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func SignUp(env *models.Env, w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: 400}
		handlers.SendResponse(w, response)
		return
	}
	//! Generate UUID
	newUUID, err := exec.Command("uuidgen").Output()
	user.Uuid = string(newUUID)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: 500}
		handlers.SendResponse(w, response)
		return
	}
	//! Generate encrypted password
	encryptedPassword, err := Encrypt(user.PasswordHash, seeders.PasswordHashingSecret)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: 500}
		handlers.SendResponse(w, response)
		return
	}
	user.PasswordHash = encryptedPassword
	//! Create user
	if err := env.Create(&user).Error; err != nil {
		response := models.Response{Message: err.Error(), Status: 500}
		handlers.SendResponse(w, response)
		return
	}
	//! Genrate token
	tokenString, tokenError := apigateway.CreateToken(string(newUUID), user.IsAdmin)
	if tokenError != nil {
		response := models.Response{Message: "Failed to create token", Status: 500}
		handlers.SendResponse(w, response)
		return
	}
	response := models.Response{Message: "Created new user", Status: 200, Jwt: &tokenString}
	handlers.SendResponse(w, response)
}
