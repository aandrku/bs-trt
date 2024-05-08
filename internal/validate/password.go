package validate

import (
	"errors"
	"unicode"
)

func Password(password string) error {
	switch {
	case len(password) < 8:
		return errors.New("Password should be at least 8 characters long")
	case len(password) > 64:
		return errors.New("Password should not exceed 64 characters")
	case !contains(password, unicode.IsUpper):
		return errors.New("Password should contain at least one upper case letter")
	case !contains(password, unicode.IsLower):
		return errors.New("Password should contain at least one lower case letter")
	case !contains(password, unicode.IsDigit):
		return errors.New("Password should contain at least one digit")
	default:
		return nil
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
