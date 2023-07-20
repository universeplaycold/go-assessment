package pkg

import (
	"errors"
	"regexp"
)

func Validate(u User) error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	matched, _ := regexp.MatchString(`^[\w-]+@([\w-]+\.)+[\w-]{2,4}$`, u.Email)
	if !matched {
		return errors.New("invalid email format")
	}

	return nil
}
