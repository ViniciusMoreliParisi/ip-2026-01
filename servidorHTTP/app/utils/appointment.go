package utils

import "log"

type Appointment struct {
	ID               int
	PatientID        int
	PatientName      string
	PatientCPF       string
	PatientTelephone string
	PatientBloodType string
	PatientAllergies string
	PatientDiseases  string
	PatientMedicines string
	AppointmentDate  string
	AppointmentTime  string
	PaymentMethod    string
	PaymentLabel     string
}

var paymentLabels = map[string]string{
	"cash":      "Dinheiro",
	"card":      "Cartão",
	"pix":       "PIX",
	"insurance": "Convênio",
}

func PaymentLabel(method string) string {
	if label, ok := paymentLabels[method]; ok {
		return label
	}
	return method
}

func InsertAppointment(patientID int, date, timeValue, paymentMethod string) error {
	query := `
		INSERT INTO appointments (patient_id, appointment_date, appointment_time, payment_method)
		VALUES ($1, $2, $3, $4)`
	_, err := DB.Exec(query, patientID, date, timeValue, paymentMethod)
	if err != nil {
		log.Printf("erro ao inserir consulta: %v", err)
		return err
	}
	return nil
}

const appointmentSelectQuery = `
		SELECT
			a.id,
			a.patient_id,
			COALESCE(p.name, ''),
			COALESCE(p.cpf, ''),
			COALESCE(p.telephone, ''),
			COALESCE(p.blood_type, ''),
			COALESCE(p.allergies, ''),
			COALESCE(p.diseases, ''),
			COALESCE(p.medicine_usage, ''),
			TO_CHAR(a.appointment_date, 'YYYY-MM-DD'),
			TO_CHAR(a.appointment_time, 'HH24:MI'),
			a.payment_method
		FROM appointments a
		INNER JOIN patients p ON p.id = a.patient_id`

func scanAppointments(rows interface {
	Next() bool
	Scan(dest ...any) error
}) ([]Appointment, error) {
	var appointments []Appointment
	for rows.Next() {
		var a Appointment
		if err := rows.Scan(
			&a.ID,
			&a.PatientID,
			&a.PatientName,
			&a.PatientCPF,
			&a.PatientTelephone,
			&a.PatientBloodType,
			&a.PatientAllergies,
			&a.PatientDiseases,
			&a.PatientMedicines,
			&a.AppointmentDate,
			&a.AppointmentTime,
			&a.PaymentMethod,
		); err != nil {
			return nil, err
		}
		a.PatientCPF = FormatCPF(a.PatientCPF)
		a.PaymentLabel = PaymentLabel(a.PaymentMethod)
		appointments = append(appointments, a)
	}
	return appointments, nil
}


func GetAppointmentsInMonth(year, month int) ([]Appointment, error) {
	start, end := MonthBounds(year, month)
	query := appointmentSelectQuery + `
		WHERE a.appointment_date >= $1
		  AND a.appointment_date <= $2
		ORDER BY a.appointment_date, a.appointment_time`

	rows, err := DB.Query(query, start.Format("2006-01-02"), end.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments, err := scanAppointments(rows)
	if err != nil {
		return nil, err
	}
	return appointments, rows.Err()
}

func IsTimeSlotAvailable(date, timeValue string) (bool, error) {
	query := `SELECT COUNT(*) FROM appointments WHERE appointment_date = $1 AND appointment_time = $2`
	var count int
	err := DB.QueryRow(query, date, timeValue).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func CountAppointmentsByDate(appointments []Appointment) map[string]int {
	counts := make(map[string]int)
	for _, a := range appointments {
		counts[a.AppointmentDate]++
	}
	return counts
}

func FilterAppointmentsByDate(appointments []Appointment, date string) []Appointment {
	var filtered []Appointment
	for _, a := range appointments {
		if a.AppointmentDate == date {
			filtered = append(filtered, a)
		}
	}
	return filtered
}
