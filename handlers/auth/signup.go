package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/ArjunMalhotra07/gorm_recruiter/bootstrap"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
//	@Summary		Sign up a user
//	@Description	Signs up a user, stores their details, and sends a welcome email
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User		true	"User signup details"
//	@Success		200		{object}	models.Response	"User successfully created"
//	@Failure		400		{object}	models.Response	"Invalid request body or email format"
//	@Failure		500		{object}	models.Response	"Internal server error"
//	@Router			/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	//! Decode incoming json body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusBadRequest, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}
	//! Validate email
	if user.Email == "" {
		response := models.Response{Message: "Email is required"}
		c.JSON(http.StatusBadRequest, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}

	//! Check if email is valid
	if _, err := mail.ParseAddress(user.Email); err != nil {
		response := models.Response{Message: "Invalid email format"}
		c.JSON(http.StatusBadRequest, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}
	//! Generate UUID
	newUUID := h.repo.CreateUserID()
	user.UserID = newUUID
	//! Generate encrypted password
	encryptedPassword, err := h.repo.CreateEncryptedPassword(user.PasswordHash, seeders.PasswordHashingSecret)
	if err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}
	user.PasswordHash = encryptedPassword
	//! Create user
	if err := h.repo.CreateUser(&user); err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}
	//! Generate token
	tokenString, tokenError := h.repo.CreateJwtToken(newUUID, user.IsEmployer)
	if tokenError != nil {
		response := models.Response{Message: "Failed to create token"}
		c.JSON(http.StatusInternalServerError, response)
		bootstrap.UserSignups.WithLabelValues("failure").Inc()
		return
	}
	bootstrap.UserSignups.WithLabelValues("success").Inc()
	response := models.Response{Message: "Created new user", Jwt: &tokenString}
	c.JSON(http.StatusOK, response)
	//! Send welcome email
	body := fmt.Sprintf("Hey %s, You have successfully signed up! Your profile headline `%s` is super catchy. We hope to provide you better services at %s", user.Name, user.ProfileHeadline, user.Address)
	emailErr := h.repo.SendWelcomeEmail(user.Email, "Welcome to Our Platform!", body)
	if emailErr != nil {
		log.Printf("Failed to send email: %v", emailErr)
	}
}
