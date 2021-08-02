package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname is already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email is already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title is already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
