package validate

import (
	"errors"
	"unicode"
)

func Password(password string) (bool, error) {
	switch {
	case len(password) < 8:
		return false, errors.New("Password should be at least 8 characters long")
	case len(password) > 64:
		return false, errors.New("Password should not exceed 64 characters")
	case !contains(password, unicode.IsUpper):
		return false, errors.New("Password should contain at least one upper case letter")
	case !contains(password, unicode.IsLower):
		return false, errors.New("Password should contain at least one lower case letter")
	case !contains(password, unicode.IsDigit):
		return false, errors.New("Password should contain at least one digit")
	default:
		return true, nil
	}
}

func contains(password string, checker func(rune) bool) bool {
	res := false

	for _, char := range password {
		if checker(char) {
			res = true
			break
		}	
	}

	return res
}
