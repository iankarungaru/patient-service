package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/iankarungaru/patient-service/models"
)

var patients = []models.Patient{
	{ID: 1, FirstName: "Alice", LastName: "Wanjiru", Phone: "0712345678"},
	{ID: 2, FirstName: "Brian", LastName: "Otieno", Phone: "0723456789"},
	{ID: 3, FirstName: "Grace", LastName: "Mwikali", Phone: "0734567890"},
}

func PatientsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(patients)
		return
	}
	if r.Method == http.MethodPost {
		createPatientHandler(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func createPatientHandler(w http.ResponseWriter, r *http.Request) {
	var newPatient models.Patient
	err := json.NewDecoder(r.Body).Decode(&newPatient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	patients = append(patients, newPatient)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPatient)
}

func PatientDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getPatientByID(w, id)
	case http.MethodPut:
		updatePatient(w, r, id)
	case http.MethodDelete:
		deletePatient(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPatientByID(w http.ResponseWriter, id int) {
	for _, p := range patients {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}

func updatePatient(w http.ResponseWriter, r *http.Request, id int) {
	var updated models.Patient
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for i, p := range patients {
		if p.ID == id {
			updated.ID = id
			patients[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}

func deletePatient(w http.ResponseWriter, id int) {
	for i, p := range patients {
		if p.ID == id {
			patients = append(patients[:i], patients[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Patient not found", http.StatusNotFound)
}
