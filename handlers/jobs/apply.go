package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *JobsHandler) ApplyToJob(c *gin.Context) {
	jobID := c.Param(constants.JobID)
	claimsInterface, _ := c.Get(constants.Claims)
	claims, _ := claimsInterface.(jwt.MapClaims)
	userID, _ := claims[constants.UniqueID].(string)
	//! Check if job exists
	err := h.repo.CheckIfJobExists(jobID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Job doesn't exist or has been deleted!"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := models.Response{Message: "Error checking job existence"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//! Check if user has already applied
	existingApplication, err := h.repo.CheckIfApplied(userID, jobID)
	if err == nil {
		response := models.Response{Message: "You have already applied for this job!", Data: existingApplication}
		c.JSON(http.StatusConflict, response)
		return
	} else if err != gorm.ErrRecordNotFound {
		response := models.Response{Message: "Error checking application status"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//! Generate Application ID
	applicationID := uuid.New().String()
	//! Add application to table
	err = h.repo.CreateApplication(applicationID, userID, jobID)
	if err != nil {
		response := models.Response{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{Message: "Job Applied successfully!"}
	c.JSON(http.StatusOK, response)
}
