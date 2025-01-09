package handlers

import repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/auth"

type AuthHandler struct {
	repo repo.AuthRepo
}

func NewAuthHandler(repository repo.AuthRepo) *AuthHandler {
	return &AuthHandler{repo: repository}
}
