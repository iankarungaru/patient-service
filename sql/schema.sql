CREATE TABLE patients (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	gender VARCHAR(10),
	date_of_birth DATE NOT NULL,
	national_id VARCHAR(20) UNIQUE,
	phone VARCHAR(20) NOT NULL,
	email VARCHAR(150),
	address TEXT,
	blood_group VARCHAR(5),
	allergies TEXT,
	emergency_contact_name VARCHAR(150),
	emergency_contact_phone VARCHAR(20),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);