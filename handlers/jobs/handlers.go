package handlers

import repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/jobs"

type JobsHandler struct {
	repo *repo.JobRepo
}

func NewJobHandler(repo *repo.JobRepo) *JobsHandler {
	return &JobsHandler{repo: repo}
}
