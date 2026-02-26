package models

import (
	"time"
)

type ApplicationStatus string

const (
	StatusApplied   ApplicationStatus = "applied"
	StatusInterview ApplicationStatus = "interview"
	StatusOffer     ApplicationStatus = "offer"
	StatusRejected  ApplicationStatus = "rejected"
	StatusAccepted  ApplicationStatus = "accepted"
)

type Application struct {
	ID        int64
	VacancyID int64
	AppliedAt time.Time
	Status    ApplicationStatus
}

var allowedTransitions = map[ApplicationStatus][]ApplicationStatus{
	StatusApplied:   {StatusInterview, StatusRejected},
	StatusInterview: {StatusOffer, StatusRejected},
	StatusOffer:     {StatusAccepted, StatusRejected},
	StatusAccepted:  {}, //final status
	StatusRejected:  {}, //final status
}

func IsValidStatusTransition(from, to ApplicationStatus) bool {
	nextStatuses, ok := allowedTransitions[from]
	if !ok {
		return false
	}

	for _, status := range nextStatuses {
		if status == to {
			return true
		}
	}
	return false
}
