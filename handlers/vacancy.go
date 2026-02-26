package handlers

import (
	"encoding/json"
	"net/http"

	"job-application-tracker-api/models"
)

type CreateVacancyRequest struct {
	CompanyName string  `json:"company_name"`
	Position    string  `json:"position"`
	SalaryMin   float64 `json:"salary_min"`
	SalaryMax   float64 `json:"salary_max"`
	Source      string  `json:"source"`
	Notes       string  `json:"notes"`
	VacancyURL  string  `json:"vacancy_url"`
}

var (
	vacancies []models.Vacancy
	nextID    int64 = 1
)

func CreateVacancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateVacancyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.CompanyName == "" {
		http.Error(w, "company_name is required", http.StatusBadRequest)
		return
	}

	if req.Position == "" {
		http.Error(w, "position is required", http.StatusBadRequest)
		return
	}

	vacancy := models.Vacancy{
		ID:          nextID,
		CompanyName: req.CompanyName,
		Position:    req.Position,
		SalaryMin:   req.SalaryMin,
		SalaryMax:   req.SalaryMax,
		Source:      req.Source,
		Notes:       req.Notes,
		VacancyURL:  req.VacancyURL,
	}

	nextID++
	vacancies = append(vacancies, vacancy)

	w.WriteHeader(http.StatusCreated)
}

func GetVacancies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vacancies); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
