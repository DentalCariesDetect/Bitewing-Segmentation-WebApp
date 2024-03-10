package errors

type PatientError struct {
	ServerInternalError  error
	PatientNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type PatientNotFoundError struct{}

func (e *PatientNotFoundError) Error() string {
	return "Patient not found"
}
