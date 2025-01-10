package handlers

import repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/employer"

type EmployerHandler struct {
	repo *repo.EmployerRepo
}

func NewEmployerHandler(repository *repo.EmployerRepo) *EmployerHandler {
	return &EmployerHandler{repo: repository}
}
