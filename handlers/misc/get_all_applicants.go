package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/gin-gonic/gin"
)

func (h *MiscHandler) GetAllApplicants(c *gin.Context) {
	//! Fetch data from DB
	applicants, err := h.repo.GetAllApplicants()
	if err != nil {
		response := models.Response{Message: "Error fetching applicants"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := models.Response{
		Message: "Applicants fetched successfully!",
		Data:    applicants,
	}
	c.JSON(http.StatusOK, response)
}
