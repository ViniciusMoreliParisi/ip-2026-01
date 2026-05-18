-- Health Management — MySQL schema

CREATE TABLE IF NOT EXISTS patients (
    id             INT AUTO_INCREMENT PRIMARY KEY,
    name           VARCHAR(150) NOT NULL,
    cpf            VARCHAR(11)  NOT NULL UNIQUE,
    telephone      VARCHAR(20)  NOT NULL,
    address        TEXT         NOT NULL,
    blood_type     VARCHAR(5)   NOT NULL,
    allergies      TEXT,
    diseases       TEXT,
    medicine_usage TEXT,
    created_at     TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_blood_type
        CHECK (blood_type IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'))
);

CREATE TABLE IF NOT EXISTS appointments (
    id               INT AUTO_INCREMENT PRIMARY KEY,
    patient_id       INT          NOT NULL,
    appointment_date DATE         NOT NULL,
    appointment_time TIME         NOT NULL,
    payment_method   VARCHAR(50)  NOT NULL,
    created_at       TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_appointments_patient
        FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE,
    CONSTRAINT chk_payment_method
        CHECK (payment_method IN ('cash', 'card', 'pix', 'insurance'))
);

CREATE INDEX idx_appointments_date ON appointments (appointment_date, appointment_time);
CREATE INDEX idx_appointments_patient ON appointments (patient_id);
