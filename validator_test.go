package validator

import (
	"testing"
)

func TestIsPhone(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"+254722333444", true},
		{"5103805566", true},
		{"3456", false},
		{"B343478", false},
		{"", false},
	}
	for _, test := range tests {
		actual := IsPhone(test.input)
		if actual != test.expected {
			t.Errorf("IsPhone(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsIP(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"255.255,0", false},
		{"234.54.12.4", true},
		{"192.168.1.1", true},
		{"", false},
	}
	for _, test := range tests {
		actual := IsIP(test.input)
		if actual != test.expected {
			t.Errorf("IsIP(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"a@bc.co", true},
		{"d+1@gmail.com", true},
		{"zatiti.com", false},
		{"e.f@gh.me", true},
		{"", false},
	}
	for _, test := range tests {
		actual := IsEmail(test.input)
		if actual != test.expected {
			t.Errorf("IsEmail(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsDomain(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"www.bc.co", true},
		{"ke.gmail.com", true},
		{"www.zatiti.com", true},
		{"a.gh.me", true},
		{"b.kat.ph", true},
		{"ke", false},
		{"", false},
	}
	for _, test := range tests {
		actual := IsDomain(test.input)
		if actual != test.expected {
			t.Errorf("IsDomain(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ff91be32-5908-452c-b4cd-8b975ac111f2", true},
		{"abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef", false},
		{"", false},
	}
	for _, test := range tests {
		actual := IsUUID(test.input)
		if actual != test.expected {
			t.Errorf("IsUUID(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}
