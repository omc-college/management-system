package validate

import "fmt"

var ErrNoSymbols = fmt.Errorf("field must not be empty")
var ErrToMuchSymbols = fmt.Errorf("to much symbols")
var ErrInvalidEmail = fmt.Errorf("invalid characters in the email")
var ErrEmailExists = fmt.Errorf("email already exists")
var ErrInvalidUserLoginInformation = fmt.Errorf("invalid email or passsword")
