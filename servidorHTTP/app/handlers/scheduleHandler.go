package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"servidorHTTP/app/utils"
)

type schedulePageData struct {
	Calendar          utils.CalendarView
	SelectedDate      string
	SelectedDateLabel string
	DayAppointments   []utils.Appointment
	TotalUpcoming     int
	HasSelection      bool
}

var scheduleTmpl = template.Must(template.ParseFiles("app/templates/schedule.html"))

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	year, month := utils.ParseYearMonth(r.URL.Query().Get("year"), r.URL.Query().Get("month"))
	selectedDate := strings.TrimSpace(r.URL.Query().Get("date"))

	appointments, err := utils.GetAppointmentsInMonth(year, month)
	if err != nil {
		http.Error(w, "Erro ao carregar agenda", http.StatusInternalServerError)
		return
	}
	if appointments == nil {
		appointments = []utils.Appointment{}
	}

	totalUpcoming := 0
	nowStr := time.Now().Format("2006-01-02 15:04")
	for _, a := range appointments {
		if a.AppointmentDate+" "+a.AppointmentTime >= nowStr {
			totalUpcoming++
		}
	}

	counts := utils.CountAppointmentsByDate(appointments)

	if selectedDate == "" {
		selectedDate = defaultSelectedDate(appointments)
	}

	calendar := utils.BuildCalendar(year, month, selectedDate, counts)
	dayAppointments := utils.FilterAppointmentsByDate(appointments, selectedDate)

	data := schedulePageData{
		Calendar:          calendar,
		SelectedDate:      selectedDate,
		SelectedDateLabel: utils.FormatDateBR(selectedDate),
		DayAppointments:   dayAppointments,
		TotalUpcoming:     totalUpcoming,
		HasSelection:      selectedDate != "",
	}

	if dayAppointments == nil {
		data.DayAppointments = []utils.Appointment{}
	}

	if err := scheduleTmpl.Execute(w, data); err != nil {
		log.Printf("erro ao renderizar agenda: %v", err)
	}
}

func defaultSelectedDate(appointments []utils.Appointment) string {
	today := time.Now().Format("2006-01-02")
	for _, a := range appointments {
		if a.AppointmentDate == today {
			return today
		}
	}
	if len(appointments) > 0 {
		return appointments[0].AppointmentDate
	}
	return today
}
