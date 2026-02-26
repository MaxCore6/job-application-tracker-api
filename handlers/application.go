package handlers

import (
	"encoding/json"
	"job-application-tracker-api/models"
	"job-application-tracker-api/service"
	"net/http"
	"strconv"
	"strings"
)

type CreateApplicationRequest struct {
	VacancyID int64 `json:"vacancy_id"`
}

type UpdateApplicationStatusRequest struct {
	Status models.ApplicationStatus `json:"status"`
}

type ApplicationHandler struct {
	service *service.ApplicationService
}

func NewApplicationHandler(s *service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{service: s}
}

func (h *ApplicationHandler) UpdateApplicationStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/applications/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid application id", http.StatusBadRequest)
		return
	}

	var req UpdateApplicationStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateStatus(id, req.Status)
	if err != nil {
		switch err {
		case service.ErrApplicationNotFound:
			http.Error(w, "application not found", http.StatusNotFound)
		case service.ErrInvalidStatusTransition:
			http.Error(w, "invalid status transition", http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ApplicationHandler) CreateApplication(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateApplicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.VacancyID == 0 {
		http.Error(w, "vacancy_id is requierd", http.StatusBadRequest)
		return
	}

	app, err := h.service.CreateApplication(req.VacancyID)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(app)
}
