package validation

import "fmt"

var ErrNoSymbols = fmt.Errorf("field must not be empty")
var ErrToMuchSymbols = fmt.Errorf("to much symbols")
var ErrInvalidEmail = fmt.Errorf("invalid email")
var ErrEmailExists = fmt.Errorf("email already exists")
var ErrConflict = fmt.Errorf("new password and old identical")
