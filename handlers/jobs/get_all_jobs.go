package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *JobsHandler) GetAllJobs(c *gin.Context) {
	//! Fetch data from DB
	jobs, err := h.repo.GetAllJobs()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response := models.Response{Message: "No Active Jobs!"}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := models.Response{Message: "Error fetching job details"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{
		Message: "Jobs fetched successfully!",
		Data:    jobs,
	}
	c.JSON(http.StatusOK, response)
}
