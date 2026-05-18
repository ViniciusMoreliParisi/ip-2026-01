package utils

import (
	"regexp"
	"strings"
	"time"
)

var phoneDigits = regexp.MustCompile(`\D`)

var validPaymentMethods = map[string]bool{
	"cash":      true,
	"card":      true,
	"pix":       true,
	"insurance": true,
}

var validBloodTypes = map[string]bool{
	"A+":  true,
	"A-":  true,
	"B+":  true,
	"B-":  true,
	"AB+": true,
	"AB-": true,
	"O+":  true,
	"O-":  true,
}

// NormalizePhone keeps only digits from the input string.
func NormalizePhone(phone string) string {
	return phoneDigits.ReplaceAllString(phone, "")
}

// NormalizeCPF keeps only digits from a CPF string.
func NormalizeCPF(cpf string) string {
	return phoneDigits.ReplaceAllString(cpf, "")
}

// FormatCPF returns CPF as 000.000.000-00.
func FormatCPF(cpf string) string {
	digits := NormalizeCPF(cpf)
	if len(digits) != 11 {
		return cpf
	}
	return digits[0:3] + "." + digits[3:6] + "." + digits[6:9] + "-" + digits[9:11]
}

// ValidateCPF checks CPF digits using the official verifier algorithm.
func ValidateCPF(cpf string) bool {
	digits := NormalizeCPF(cpf)
	if len(digits) != 11 {
		return false
	}

	allEqual := true
	for i := 1; i < 11; i++ {
		if digits[i] != digits[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(digits[i]-'0') * (10 - i)
	}
	remainder := sum % 11
	firstDigit := 0
	if remainder >= 2 {
		firstDigit = 11 - remainder
	}
	if int(digits[9]-'0') != firstDigit {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(digits[i]-'0') * (11 - i)
	}
	remainder = sum % 11
	secondDigit := 0
	if remainder >= 2 {
		secondDigit = 11 - remainder
	}
	return int(digits[10]-'0') == secondDigit
}

// ValidateBloodType returns true for allowed blood type values.
func ValidateBloodType(bloodType string) bool {
	return validBloodTypes[strings.TrimSpace(bloodType)]
}

// ValidatePhone checks Brazilian-style numbers (10 or 11 digits).
func ValidatePhone(phone string) bool {
	digits := NormalizePhone(phone)
	return len(digits) == 10 || len(digits) == 11
}

// ValidatePaymentMethod returns true for allowed payment values.
func ValidatePaymentMethod(method string) bool {
	return validPaymentMethods[strings.ToLower(strings.TrimSpace(method))]
}

// ValidateAppointmentDateTime ensures the appointment is not in the past.
func ValidateAppointmentDateTime(dateStr, timeStr string) (time.Time, bool) {
	dateStr = strings.TrimSpace(dateStr)
	timeStr = strings.TrimSpace(timeStr)
	if dateStr == "" || timeStr == "" {
		return time.Time{}, false
	}

	appointmentAt, err := time.ParseInLocation("2006-01-02 15:04", dateStr+" "+timeStr, time.Local)
	if err != nil {
		return time.Time{}, false
	}

	now := time.Now().Truncate(time.Minute)
	return appointmentAt, !appointmentAt.Before(now)
}
