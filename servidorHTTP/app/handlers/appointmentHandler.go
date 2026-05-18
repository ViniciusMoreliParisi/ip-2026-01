package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"servidorHTTP/app/utils"
)

type appointmentFormData struct {
	Patients   []utils.Patient
	MinDate    string
	Error      string
	Success    bool
}

var appointmentFormTmpl = template.Must(
	template.New("appointment_form").Funcs(template.FuncMap{
		"formatCPF": utils.FormatCPF,
	}).ParseFiles("app/templates/appointment_form.html"),
)

func AppointmentFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	patients, err := utils.GetAllPatients()
	if err != nil {
		http.Error(w, "Erro ao carregar pacientes", http.StatusInternalServerError)
		return
	}

	data := appointmentFormData{
		Patients: patients,
		MinDate:  time.Now().Format("2006-01-02"),
		Error:    r.URL.Query().Get("error"),
		Success:  r.URL.Query().Get("success") == "1",
	}

	if err := appointmentFormTmpl.ExecuteTemplate(w, "appointment_form.html", data); err != nil {
		log.Printf("erro ao renderizar formulário de consulta: %v", err)
	}
}

func AppointmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	patientID, err := strconv.Atoi(strings.TrimSpace(r.FormValue("patient_id")))
	if err != nil || patientID <= 0 {
		http.Redirect(w, r, "/appointments/new?error=patient", http.StatusSeeOther)
		return
	}

	date := strings.TrimSpace(r.FormValue("appointment_date"))
	timeValue := strings.TrimSpace(r.FormValue("appointment_time"))
	if len(timeValue) > 5 {
		timeValue = timeValue[:5]
	}
	paymentMethod := strings.ToLower(strings.TrimSpace(r.FormValue("payment_method")))

	if !utils.ValidatePaymentMethod(paymentMethod) {
		http.Redirect(w, r, "/appointments/new?error=payment", http.StatusSeeOther)
		return
	}

	if _, valid := utils.ValidateAppointmentDateTime(date, timeValue); !valid {
		http.Redirect(w, r, "/appointments/new?error=datetime", http.StatusSeeOther)
		return
	}

	available, err := utils.IsTimeSlotAvailable(date, timeValue)
	if err != nil {
		http.Error(w, "Erro ao verificar disponibilidade", http.StatusInternalServerError)
		return
	}
	if !available {
		http.Redirect(w, r, "/appointments/new?error=conflict", http.StatusSeeOther)
		return
	}

	exists, err := utils.PatientExists(patientID)
	if err != nil {
		http.Error(w, "Erro ao validar paciente", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Redirect(w, r, "/appointments/new?error=patient", http.StatusSeeOther)
		return
	}

	if err := utils.InsertAppointment(patientID, date, timeValue, paymentMethod); err != nil {
		http.Error(w, "Erro ao agendar consulta", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/appointments/new?success=1", http.StatusSeeOther)
}
