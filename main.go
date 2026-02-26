package main

import (
	"log"
	"net/http"

	"job-application-tracker-api/handlers"
	"job-application-tracker-api/repo"
	"job-application-tracker-api/service"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/vacancies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateVacancy(w, r)
		case http.MethodGet:
			handlers.GetVacancies(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	appRepo := repo.NewApplicationRepository()
	appService := service.NewApplicationService(appRepo)
	appHandler := handlers.NewApplicationHandler(appService)

	mux.HandleFunc("/applications", appHandler.CreateApplication)
	mux.HandleFunc("/applications/", appHandler.UpdateApplicationStatus)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
