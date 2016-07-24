package validator

import (
	"regexp"
)

func IsNotEmpty(s string) bool {
	return len(s) > 0
}

func IsInteger(s string) bool {
	re := regexp.MustCompile(`^-{0,1}\d+$`)
	return re.MatchString(s)
}

func IsPositiveInteger(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func IsNegativeInteger(s string) bool {
	re := regexp.MustCompile(`^-\d+$`)
	return re.MatchString(s)
}

func IsCurrency(s string) bool {
	re := regexp.MustCompile(`^-{0,1}\d*\.{0,2}\d+$`)
	return re.MatchString(s)
}

func IsDouble(s string) bool {
	re := regexp.MustCompile(`^-{0,1}\d*\.{0,1}\d+$`)
	return re.MatchString(s)
}

func IsPositiveDouble(s string) bool {
	re := regexp.MustCompile(`^\d*\.{0,1}\d+$`)
	return re.MatchString(s)
}

func IsNegativeDouble(s string) bool {
	re := regexp.MustCompile(`^-\d*\.{0,1}\d+$`)
	return re.MatchString(s)
}

func IsPhone(s string) bool {
	re := regexp.MustCompile(`^\+?[0-9\s]{8,16}`)
	return re.MatchString(s)
}

func IsYear(s string) bool {
	re := regexp.MustCompile(`^(19|20)[\d]{2,2}$`)
	return re.MatchString(s)
}

func IsIP(s string) bool {
	re := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	return re.MatchString(s)
}

func IsEmail(s string) bool {
	re := regexp.MustCompile(`^([a-z0-9_\.\+-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	return re.MatchString(s)
}

func IsDomain(s string) bool {
	re := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`)
	return re.MatchString(s)
}

func IsUUID(s string) bool {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[4][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return re.MatchString(s)
}
