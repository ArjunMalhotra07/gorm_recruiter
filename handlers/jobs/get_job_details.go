package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *JobsHandler) GetJobData(c *gin.Context) {
	//! Get Job ID
	jobID := c.Param(constants.JobID)
	//! Fetch data from DB
	job, err := h.repo.GetJobData(jobID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "Job doesn't exist or is either deleted!"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := models.Response{Message: "Error fetching job details"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{
		Message: "Job fetched successfully!",
		Data:    job,
	}
	c.JSON(http.StatusOK, response)
}
