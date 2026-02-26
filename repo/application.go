package repo

import (
	"errors"
	"job-application-tracker-api/models"
	"time"
)

var ErrNotFound = errors.New("not found")

type ApplicationRepository struct {
	applications []models.Application
	nextID       int64
}

func NewApplicationRepository() *ApplicationRepository {
	return &ApplicationRepository{
		applications: []models.Application{},
		nextID:       1,
	}
}

func (r *ApplicationRepository) GetByID(id int64) (*models.Application, error) {
	for i := range r.applications {
		if r.applications[i].ID == id {
			return &r.applications[i], nil
		}
	}
	return nil, ErrNotFound
}

func (r *ApplicationRepository) UpdateStatus(id int64, status models.ApplicationStatus) error {
	for i := range r.applications {
		if r.applications[i].ID == id {
			r.applications[i].Status = status
			return nil
		}
	}
	return ErrNotFound
}

func (r *ApplicationRepository) Create(vacancyID int64) (*models.Application, error) {
	app := models.Application{
		ID:        r.nextID,
		VacancyID: vacancyID,
		AppliedAt: time.Now(),
		Status:    models.StatusApplied,
	}

	r.nextID++
	r.applications = append(r.applications, app)
	return &app, nil
}
