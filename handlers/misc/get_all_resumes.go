package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
)

func (h *MiscHandler) GetAllResumes(c *gin.Context) {
	//! Fetch data from DB
	resumes, err := h.repo.GetAllResumes()
	if err != nil {
		response := models.Response{Message: "Error fetching resumes"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{
		Message: "Resumes fetched successfully!",
		Data:    resumes,
	}
	c.JSON(http.StatusOK, response)
}
