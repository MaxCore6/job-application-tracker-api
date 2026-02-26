package service

import (
	"errors"
	"job-application-tracker-api/models"
	"job-application-tracker-api/repo"
)

var (
	ErrApplicationNotFound     = errors.New("application not found")
	ErrInvalidStatusTransition = errors.New("invalid status transition")
)

// type ApplicationRepository interface {
// 	GetByID(id int64) (*models.Application, error)
// 	UpdateStatus(id int64, status models.ApplicationStatus) error
// }

type ApplicationService struct {
	repo *repo.ApplicationRepository
}

func NewApplicationService(r *repo.ApplicationRepository) *ApplicationService {
	return &ApplicationService{repo: r}
}

func (s *ApplicationService) UpdateStatus(id int64, newStatus models.ApplicationStatus) error {
	app, err := s.repo.GetByID(id)
	if err != nil {
		return ErrApplicationNotFound
	}

	if !models.IsValidStatusTransition(app.Status, newStatus) {
		return ErrInvalidStatusTransition
	}

	return s.repo.UpdateStatus(id, newStatus)
}

func (s *ApplicationService) CreateApplication(vacancyID int64) (*models.Application, error) {
	return s.repo.Create(vacancyID)
}
