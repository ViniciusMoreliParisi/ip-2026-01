package utils

import (
	"testing"
	"time"
)

func TestValidatePhone(t *testing.T) {
	cases := []struct {
		phone string
		valid bool
	}{
		{"11987654321", true},
		{"(11) 98765-4321", true},
		{"123", false},
		{"", false},
	}

	for _, c := range cases {
		if got := ValidatePhone(c.phone); got != c.valid {
			t.Errorf("ValidatePhone(%q) = %v, want %v", c.phone, got, c.valid)
		}
	}
}

func TestValidateCPF(t *testing.T) {
	cases := []struct {
		cpf   string
		valid bool
	}{
		{"529.982.247-25", true},
		{"52998224725", true},
		{"123.456.789-09", true},
		{"111.444.777-35", true},
		{"111.111.111-11", false},
		{"000.000.000-00", false},
		{"123", false},
	}

	for _, c := range cases {
		if got := ValidateCPF(c.cpf); got != c.valid {
			t.Errorf("ValidateCPF(%q) = %v, want %v", c.cpf, got, c.valid)
		}
	}
}

func TestValidateBloodType(t *testing.T) {
	if !ValidateBloodType("O+") {
		t.Error("expected O+ to be valid")
	}
	if ValidateBloodType("X+") {
		t.Error("expected X+ to be invalid")
	}
}

func TestValidateAppointmentDateTime(t *testing.T) {
	future := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	past := time.Now().Add(-48 * time.Hour).Format("2006-01-02")

	if _, ok := ValidateAppointmentDateTime(future, "10:00"); !ok {
		t.Error("expected future appointment to be valid")
	}

	if _, ok := ValidateAppointmentDateTime(past, "10:00"); ok {
		t.Error("expected past appointment to be invalid")
	}
}
