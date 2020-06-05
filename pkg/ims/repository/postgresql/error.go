package postgresql

<<<<<<< HEAD
import "fmt"

var ErrNoRows = fmt.Errorf("no rows with such id")
var ErrConvertId = fmt.Errorf("can not convert id")
var ErrCloseStmt = fmt.Errorf("closing statement error")

const queryErrorMessage = "query error"
const scanErrorMessage = "scanning error"
=======

const UpdateCredentialsErrorMessage = "update credentials error"


>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674

type QueryError struct {
	Message string
	Err     error
}

func (err QueryError) Error() string {
	return err.Message
}

func (err QueryError) Unwrap() error {
	return err.Err
}

<<<<<<< HEAD
type ScanError struct {
	Message string
	Err     error
}

func (err ScanError) Error() string {
	return err.Message
}

func (err ScanError) Unwrap() error {
	return err.Err
}

=======
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
