package db

import (
	"context"

	"github.com/iankarungaru/patient-service/models"
)

func GetAllPatients() ([]models.Patient, error) {
	rows, err := Pool.Query(context.Background(), "SELECT id, first_name, last_name, phone FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient

	for rows.Next() {
		var p models.Patient
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}
