package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"servidorHTTP/app/utils"
)

type patientFormData struct {
	Error   string
	Success bool
}

var patientFormTmpl = template.Must(template.ParseFiles("app/templates/patient_form.html"))

func PatientFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	data := patientFormData{
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success") == "1",
	}

	if err := patientFormTmpl.Execute(w, data); err != nil {
		log.Printf("erro ao renderizar formulário de paciente: %v", err)
	}
}

func PatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao ler formulário", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	cpf := strings.TrimSpace(r.FormValue("cpf"))
	telephone := strings.TrimSpace(r.FormValue("telephone"))
	address := strings.TrimSpace(r.FormValue("address"))
	bloodType := strings.TrimSpace(r.FormValue("blood_type"))
	allergies := strings.TrimSpace(r.FormValue("allergies"))
	diseases := strings.TrimSpace(r.FormValue("diseases"))
	medicineUsage := strings.TrimSpace(r.FormValue("medicine_usage"))

	if name == "" || address == "" || cpf == "" || bloodType == "" {
		http.Redirect(w, r, "/patients/new?error=required", http.StatusSeeOther)
		return
	}

	if !utils.ValidateCPF(cpf) {
		http.Redirect(w, r, "/patients/new?error=cpf", http.StatusSeeOther)
		return
	}

	if !utils.ValidateBloodType(bloodType) {
		http.Redirect(w, r, "/patients/new?error=blood_type", http.StatusSeeOther)
		return
	}

	if !utils.ValidatePhone(telephone) {
		http.Redirect(w, r, "/patients/new?error=phone", http.StatusSeeOther)
		return
	}

	normalizedCPF := utils.NormalizeCPF(cpf)
	exists, err := utils.CPFExists(normalizedCPF)
	if err != nil {
		log.Printf("erro ao consultar CPF no banco: %v", err)
		http.Redirect(w, r, "/patients/new?error=db", http.StatusSeeOther)
		return
	}
	if exists {
		http.Redirect(w, r, "/patients/new?error=cpf_duplicate", http.StatusSeeOther)
		return
	}

	input := utils.PatientInput{
		Name:          name,
		CPF:           normalizedCPF,
		Telephone:     utils.NormalizePhone(telephone),
		Address:       address,
		BloodType:     bloodType,
		Allergies:     allergies,
		Diseases:      diseases,
		MedicineUsage: medicineUsage,
	}

	if err := utils.InsertPatient(input); err != nil {
		http.Error(w, "Erro ao salvar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/patients/new?success=1", http.StatusSeeOther)
}
