package main

import (
	"fmt"
	"net/http"
	

	"github.com/iankarungaru/patient-service/handlers"
)





func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/patients", handlers.PatientsHandler)
	mux.HandleFunc("/patients/{id}", handlers.PatientDetailHandler)
    
	fmt.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", handlers.LoggingMiddleware(mux))
}






