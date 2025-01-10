package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *EmployerHandler) GetMyJobsDetail(c *gin.Context) {
	claimsInterface, _ := c.Get(constants.Claims)
	claims, _ := claimsInterface.(jwt.MapClaims)
	userID, _ := claims[constants.UniqueID].(string)

	//! Fetch jobs posted by the user
	jobs, err := h.repo.GetJobsPostedByUser(userID, true)
	if err != nil {
		response := models.Response{Message: "Error fetching job details!", Status: http.StatusInternalServerError}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//! Fetch applicants for each job
	var jobsWithApplicants []struct {
		models.Job
		Applicants []models.User `json:"applicants"`
	}
	for _, job := range jobs {
		applicants, err := h.repo.GetApplicantsForJob(job.JobID)
		if err != nil {
			response := models.Response{Message: "Error fetching applicants!", Status: http.StatusInternalServerError}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		jobsWithApplicants = append(jobsWithApplicants, struct {
			models.Job
			Applicants []models.User `json:"applicants"`
		}{
			Job:        job,
			Applicants: applicants,
		})
	}
	//! Send response
	response := models.Response{
		Message: "My Jobs fetched successfully!",
		Status:  http.StatusOK,
		Data:    jobsWithApplicants,
	}
	c.JSON(http.StatusOK, response)
}
