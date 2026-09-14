package models

type Patient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	DOB   string `json:"dob"`
	Phone string `json:"phone"`
}