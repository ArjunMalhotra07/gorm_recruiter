package handlers

import (
	"net/http"
	"strings"

	"github.com/ArjunMalhotra07/gorm_recruiter/constants"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *JobsHandler) UploadResume(c *gin.Context) {
	claimsInterface, _ := c.Get(constants.Claims)
	claims, _ := claimsInterface.(jwt.MapClaims)
	userID, _ := claims[constants.UniqueID].(string)
	// Get file from request
	file, header, err := c.Request.FormFile("resume")
	if err != nil {
		response := models.Response{Message: "Invalid file upload"}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	defer file.Close()

	// Save file to directory
	resumeFilePath, err := h.repo.SaveResumeToDirectory(file, header)
	if err != nil {
		response := models.Response{Message: "Error saving file"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Read file content
	fileContent, err := h.repo.ReadFileContent(resumeFilePath)
	if err != nil {
		response := models.Response{Message: "Error reading file"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Parse resume
	parsedResponse, err := h.repo.ParseResume(fileContent)
	if err != nil {
		response := models.Response{Message: "Error parsing resume"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Map response to model and save to DB
	parsedResume := &models.Resume{
		ResumeID:          uuid.NewString(),
		UserID:            userID,
		ResumeFileAddress: resumeFilePath,
		Name:              parsedResponse.Name,
		Email:             parsedResponse.Email,
		Phone:             parsedResponse.Phone,
		Skills:            strings.Join(parsedResponse.Skills, ", "),
		Educations:        parsedResponse.Education,
		Experiences:       parsedResponse.Experience,
	}
	if err := h.repo.SaveParsedResumeToDB(parsedResume); err != nil {
		response := models.Response{Message: "Error saving to database"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.Response{Message: "Resume uploaded and parsed successfully", Data: parsedResume}
	c.JSON(http.StatusOK, response)
}
