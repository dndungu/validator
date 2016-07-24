package validator

import (
	"testing"
)

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Rongai", true},
		{"鎮", true},
		{"", false},
	}
	for _, test := range tests {
		actual := IsNotEmpty(test.input)
		if actual != test.expected {
			t.Errorf("IsString(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2b", false},
		{"234", true},
		{"abc", false},
	}
	for _, test := range tests {
		actual := IsInteger(test.input)
		if actual != test.expected {
			t.Errorf("IsInteger(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsPositiveInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"a4", false},
		{"123", true},
		{"-3", false},
		{"-v4", false},
	}
	for _, test := range tests {
		actual := IsPositiveInteger(test.input)
		if actual != test.expected {
			t.Errorf("IsPositiveInteger(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsNegativeInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"5t", false},
		{"123", false},
		{"-3", true},
		{"-v4", false},
	}
	for _, test := range tests {
		actual := IsNegativeInteger(test.input)
		if actual != test.expected {
			t.Errorf("IsNegativeInteger(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsCurrency(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"50", true},
		{"12.05", true},
		{"-1", true},
		{"0.01", true},
	}
	for _, test := range tests {
		actual := IsCurrency(test.input)
		if actual != test.expected {
			t.Errorf("IsCurrency(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsDouble(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"2ef", false},
		{"1.98", true},
		{"54.06", true},
		{"-1.09", true},
		{"0.01", true},
	}
	for _, test := range tests {
		actual := IsDouble(test.input)
		if actual != test.expected {
			t.Errorf("IsDouble(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsPositiveDouble(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"2ef", false},
		{"1.98", true},
		{"54.06", true},
		{"-1.09", false},
		{"0.01", true},
	}
	for _, test := range tests {
		actual := IsPositiveDouble(test.input)
		if actual != test.expected {
			t.Errorf("IsPositiveDouble(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsNegativeDouble(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"2ef", false},
		{"-1.98", true},
		{"54.06", false},
		{"-0.01", true},
	}
	for _, test := range tests {
		actual := IsNegativeDouble(test.input)
		if actual != test.expected {
			t.Errorf("IsNegativeDouble(%s): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

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

func TestIsYear(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1980", true},
		{"99", false},
		{"2017", true},
		{"", false},
	}
	for _, test := range tests {
		actual := IsYear(test.input)
		if actual != test.expected {
			t.Errorf("IsYear(%s): expected %t, actual %t", test.input, test.expected, actual)
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
