package utils

import "log"

type Patient struct {
	ID            int
	Name          string
	CPF           string
	Telephone     string
	Address       string
	BloodType     string
	Allergies     string
	Diseases      string
	MedicineUsage string
}

type PatientInput struct {
	Name          string
	CPF           string
	Telephone     string
	Address       string
	BloodType     string
	Allergies     string
	Diseases      string
	MedicineUsage string
}

func InsertPatient(p PatientInput) error {
	query := `
		INSERT INTO patients (
			name, cpf, telephone, address, blood_type,
			allergies, diseases, medicine_usage
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := DB.Exec(
		query,
		p.Name,
		p.CPF,
		p.Telephone,
		p.Address,
		p.BloodType,
		p.Allergies,
		p.Diseases,
		p.MedicineUsage,
	)
	if err != nil {
		log.Printf("erro ao inserir paciente: %v", err)
		return err
	}
	return nil
}

func CPFExists(cpf string) (bool, error) {
	var exists bool
	err := DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM patients WHERE cpf = $1)`, cpf).Scan(&exists)
	return exists, err
}

func GetAllPatients() ([]Patient, error) {
	query := `
		SELECT id,
		       COALESCE(name, ''),
		       COALESCE(cpf, ''),
		       COALESCE(telephone, ''),
		       COALESCE(address, ''),
		       COALESCE(blood_type, ''),
		       COALESCE(allergies, ''),
		       COALESCE(diseases, ''),
		       COALESCE(medicine_usage, '')
		FROM patients ORDER BY name`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.CPF,
			&p.Telephone,
			&p.Address,
			&p.BloodType,
			&p.Allergies,
			&p.Diseases,
			&p.MedicineUsage,
		); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, rows.Err()
}

func PatientExists(id int) (bool, error) {
	var exists bool
	err := DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM patients WHERE id = $1)`, id).Scan(&exists)
	return exists, err
}
