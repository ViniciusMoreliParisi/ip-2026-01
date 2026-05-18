package utils

import (
	"fmt"
	"time"
)

var monthNamesPT = []string{
	"",
	"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
	"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
}

type CalendarDay struct {
	Day            int
	Date           string
	InMonth        bool
	IsToday        bool
	HasAppointment bool
	Count          int
	IsSelected     bool
	URL            string
}

type CalendarView struct {
	Year      int
	Month     int
	Title     string
	PrevURL   string
	NextURL   string
	Weekdays  []string
	Days      []CalendarDay
}

func FormatDateBR(date string) string {
	t, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return date
	}
	return t.Format("02/01/2006")
}

func BuildCalendar(year, month int, selectedDate string, counts map[string]int) CalendarView {
	if month < 1 {
		month = 1
		year--
	}
	if month > 12 {
		month = 12
		year++
	}

	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	startPad := weekdayIndexMonday(firstOfMonth.Weekday())
	gridStart := firstOfMonth.AddDate(0, 0, -startPad)

	today := time.Now().Format("2006-01-02")
	if selectedDate == "" {
		selectedDate = today
	}

	prev := firstOfMonth.AddDate(0, -1, 0)
	next := firstOfMonth.AddDate(0, 1, 0)

	days := make([]CalendarDay, 0, 42)
	for i := 0; i < 42; i++ {
		current := gridStart.AddDate(0, 0, i)
		dateKey := current.Format("2006-01-02")
		inMonth := current.Month() == firstOfMonth.Month() && current.Year() == firstOfMonth.Year()
		count := counts[dateKey]

		days = append(days, CalendarDay{
			Day:            current.Day(),
			Date:           dateKey,
			InMonth:        inMonth,
			IsToday:        dateKey == today,
			HasAppointment: count > 0,
			Count:          count,
			IsSelected:     dateKey == selectedDate,
			URL: fmt.Sprintf(
				"/schedule?year=%d&month=%d&date=%s",
				current.Year(), int(current.Month()), dateKey,
			),
		})
	}

	_ = lastOfMonth

	return CalendarView{
		Year:     year,
		Month:    month,
		Title:    fmt.Sprintf("%s %d", monthNamesPT[month], year),
		PrevURL:  fmt.Sprintf("/schedule?year=%d&month=%d", prev.Year(), int(prev.Month())),
		NextURL:  fmt.Sprintf("/schedule?year=%d&month=%d", next.Year(), int(next.Month())),
		Weekdays: []string{"Seg", "Ter", "Qua", "Qui", "Sex", "Sáb", "Dom"},
		Days:     days,
	}
}

func weekdayIndexMonday(weekday time.Weekday) int {
	if weekday == time.Sunday {
		return 6
	}
	return int(weekday) - 1
}

func ParseYearMonth(yearStr, monthStr string) (int, int) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	if y, err := parsePositiveInt(yearStr); err == nil && y >= 2000 && y <= 2100 {
		year = y
	}
	if m, err := parsePositiveInt(monthStr); err == nil && m >= 1 && m <= 12 {
		month = m
	}
	return year, month
}

func MonthBounds(year, month int) (time.Time, time.Time) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, -1)
	return start, end
}

func parsePositiveInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}
