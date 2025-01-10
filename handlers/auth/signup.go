package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) SignUp(c *gin.Context) {
	//! Decode incoming json body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusBadRequest}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//! Validate email
	if user.Email == "" {
		response := models.Response{Message: "Email is required", Status: http.StatusBadRequest}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//! Check if email is valid
	if _, err := mail.ParseAddress(user.Email); err != nil {
		response := models.Response{Message: "Invalid email format", Status: http.StatusBadRequest}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//! Generate UUID
	newUUID, err := h.repo.CreateUserID()
	if err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	user.UserID = string(newUUID)
	//! Generate encrypted password
	encryptedPassword, err := h.repo.CreateEncryptedPassword(user.PasswordHash, seeders.PasswordHashingSecret)
	if err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	user.PasswordHash = encryptedPassword
	//! Create user
	if err := h.repo.CreateUser(&user); err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//! Genrate token
	tokenString, tokenError := h.repo.CreateJwtToken(string(newUUID), user.IsEmployer)
	if tokenError != nil {
		response := models.Response{Message: "Failed to create token", Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{Message: "Created new user", Status: 200, Jwt: &tokenString}
	c.JSON(http.StatusOK, response)
	body := fmt.Sprintf("Hey %s, You have successfully signed up! Your profile headline `%s` is super catchy. We hope to provide you better services at %s", user.Name, user.ProfileHeadline, user.Address)
	emailErr := h.repo.SendWelcomeEmail(user.Email, "Welcome to Our Platform!", body)
	if emailErr != nil {
		log.Printf("Failed to send email: %v", emailErr)
	}
}
