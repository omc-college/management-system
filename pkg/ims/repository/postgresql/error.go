package postgresql


const UpdateCredentialsErrorMessage = "update credentials error"



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

