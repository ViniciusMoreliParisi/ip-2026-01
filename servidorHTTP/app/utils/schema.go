package utils

import "log"

// EnsureHealthSchema creates or updates tables required by Health Management.
func EnsureHealthSchema() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS patients (
			id             SERIAL PRIMARY KEY,
			name           VARCHAR(150) NOT NULL,
			cpf            VARCHAR(11)  NOT NULL UNIQUE,
			telephone      VARCHAR(20)  NOT NULL,
			address        TEXT         NOT NULL,
			blood_type     VARCHAR(5)   NOT NULL
				CHECK (blood_type IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')),
			allergies      TEXT         DEFAULT '',
			diseases       TEXT         DEFAULT '',
			medicine_usage TEXT         DEFAULT '',
			created_at     TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS appointments (
			id               SERIAL PRIMARY KEY,
			patient_id       INTEGER      NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
			appointment_date DATE         NOT NULL,
			appointment_time TIME         NOT NULL,
			payment_method   VARCHAR(50)  NOT NULL
				CHECK (payment_method IN ('cash', 'card', 'pix', 'insurance')),
			created_at       TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_appointments_date
			ON appointments (appointment_date, appointment_time)`,
		`CREATE INDEX IF NOT EXISTS idx_appointments_patient
			ON appointments (patient_id)`,
	}

	for _, stmt := range stmts {
		if _, err := DB.Exec(stmt); err != nil {
			log.Printf("aviso ao aplicar schema: %v\nSQL: %s", err, stmt)
			return err
		}
	}

	return nil
}
