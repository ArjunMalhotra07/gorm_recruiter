package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *EmployerHandler) AddJob(c *gin.Context) {
	//! Get user ID from jwt claims
	claimsInterface, _ := c.Get(constants.Claims)
	claims, _ := claimsInterface.(jwt.MapClaims)
	userID, _ := claims[constants.UniqueID].(string)
	//! Decode the incoming JSON body into a Job struct
	var currentJob models.Job
	if err := c.ShouldBindJSON(&currentJob); err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//! Generating new id for job
	newUUID := uuid.New().String()
	currentJob.JobID = newUUID
	currentJob.PostedByID = userID
	//! Add job in table
	if err := h.repo.AddJob(&currentJob); err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{Message: "Job Posted successfully!"}
	c.JSON(http.StatusOK, response)
}
