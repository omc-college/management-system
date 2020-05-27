package validation

import "github.com/omc-college/management-system/pkg/ims/models"

func Data (request *models.SignupRequest) error {
	var err error

	if request.FirstName == "" {
		err = ErrNoSymbols
		return err
	} else if len(request.FirstName) > 256 {
		err = ErrToMuchSymbols
		return err
	}

	if request.LastName == "" {
		err = ErrNoSymbols
		return err
	} else if len(request.LastName) > 256 {
		err = ErrToMuchSymbols
		return err
	}

	if request.Password == "" {
		err = ErrNoSymbols
		return err
	} else if len(request.Password) > 256 {
		err = ErrToMuchSymbols
		return err
	}

	if request.Email == "" {
		err = ErrNoSymbols
		return err
	} else if len(request.Email) > 256 {
		err = ErrToMuchSymbols
		return err
	}

	if Email(request.Email) == false {
		err = ErrInvalidEmail
		return err
	}

	return nil
}