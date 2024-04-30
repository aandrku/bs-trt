package validate

import "errors"

func Username(username string) (bool, error) {
	switch {
	case len(username) < 4:
		return false, errors.New("Username is too short")
	case len(username) > 24:
		return false, errors.New("Username is too long")
	default:
		return true, nil
	}
}
