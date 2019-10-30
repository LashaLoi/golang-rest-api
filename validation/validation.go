package validation

import (
	"errors"

	"../customtypes"
)

// UserValidation ...
func UserValidation(user customtypes.User) error {
	if user.Name == "" {
		return errors.New("Name is not exist")
	} else if user.Email == "" {
		return errors.New("Email is not exist")
	} else if user.Phone == "" {
		return errors.New("Phone is not exist")
	}

	return nil
}
