package validator

import "strings"

func IsStringEmpty(str *string) bool {
	return str == nil || strings.Trim(*str, " \t\n") == ""
}

func IsLatitudeCorrect(c *float64) bool {
	if c == nil {
		return false
	}

	if *c > 90 || *c < -90 {
		return false
	}
	return true
}

func IsLongitudeCorrect(c *float64) bool {
	if c == nil {
		return false
	}

	if *c > 180 || *c < -180 {
		return false
	}
	return true
}
