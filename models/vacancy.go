package models

type Vacancy struct {
	ID          int64
	CompanyName string
	Position    string
	SalaryMin   float64
	SalaryMax   float64
	Source      string
	Notes       string
	VacancyURL  string
}
