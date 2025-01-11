package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/seeders"
)

func (r *JobRepo) SaveResumeToDirectory(userResume multipart.File, header *multipart.FileHeader) (string, error) {
	resumeDir := "./resumes"
	if err := os.MkdirAll(resumeDir, os.ModePerm); err != nil {
		return "", err
	}

	resumeFilePath := filepath.Join(resumeDir, header.Filename)
	outFile, err := os.Create(resumeFilePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, userResume); err != nil {
		return "", err
	}

	return resumeFilePath, nil
}
func (r *JobRepo) ReadFileContent(resumeFilePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(resumeFilePath)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

func (r *JobRepo) ParseResume(fileContent []byte) (*models.ResumeResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", bytes.NewReader(fileContent))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", seeders.ResumeParserKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(bodyBytes))
	}

	var parsedResponse models.ResumeResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
func (r *JobRepo) SaveParsedResumeToDB(parsedResume *models.Resume) error {
	return r.Driver.Create(parsedResume).Error
}
