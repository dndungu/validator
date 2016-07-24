package validator

import (
	"regexp"
)

// IsPhone returns true if value is in a valid phone format
func IsPhone(s string) bool {
	re := regexp.MustCompile(`^\+?[0-9\s]{8,16}`)
	return re.MatchString(s)
}

// IsIP returns true if value is a valid ip address format
func IsIP(s string) bool {
	re := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	return re.MatchString(s)
}

// IsEmail returns true if value is a valid email address format
func IsEmail(s string) bool {
	re := regexp.MustCompile(`^([a-z0-9_\.\+-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	return re.MatchString(s)
}

// IsDomain returns true if value is a valid domain name format
func IsDomain(s string) bool {
	re := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`)
	return re.MatchString(s)
}

// IsUUID returns true if string is a valid UUID version 4
func IsUUID(s string) bool {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[4][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return re.MatchString(s)
}
