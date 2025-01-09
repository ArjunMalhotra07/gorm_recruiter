package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *AuthHandler) LogIn(c *gin.Context) {
	//! Decode the incoming JSON body into a User struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response := models.Response{Message: err.Error(), Status: http.StatusBadRequest}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//! Encrypt the user's password
	encText, err := h.repo.CreateEncryptedPassword(user.PasswordHash, seeders.PasswordHashingSecret)
	if err != nil {
		response := models.Response{Message: "Error encrypting password", Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//! Check if the user exists in the database with the provided email and password
	var currentUser models.User
	if err := h.repo.LoginUser(user.Email, encText); err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Email ID or Password doesn't match", Status: http.StatusUnauthorized}
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := models.Response{Message: err.Error(), Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//! Generate a JWT token for the authenticated user
	tokenString, tokenError := h.repo.CreateJwtToken(string(currentUser.UserID), currentUser.IsEmployer)
	if tokenError != nil {
		response := models.Response{Message: "Failed to create token", Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{Message: "User exists", Status: http.StatusOK, Jwt: &tokenString}
	c.JSON(http.StatusOK, response)

}
