package models

type ResumeResponse struct {
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Skills     []string     `json:"skills"`
	Education  []Education  `json:"education"`
	Experience []Experience `json:"experience"`
}
