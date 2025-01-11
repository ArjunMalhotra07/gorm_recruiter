package handlers

import repo "github.com/ArjunMalhotra07/gorm_recruiter/repo/misc"

type MiscHandler struct {
	repo *repo.MiscRepo
}

func NewMiscHandler(repository *repo.MiscRepo) *MiscHandler {
	return &MiscHandler{repo: repository}
}
