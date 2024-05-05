package validate

import "errors"

func Email(email string) (bool, error) {
	switch {
	case len(email) < 4:
		return false, errors.New("Email is too short")
	case len(email) > 24:
		return false, errors.New("Email is too long")
	default:
		return true, nil
	}
}
