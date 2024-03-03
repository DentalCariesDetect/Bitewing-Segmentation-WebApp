package errors

type DentistError struct {
	ServerInternalError  error
	DentistNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type DentistNotFoundError struct{}

func (e *DentistNotFoundError) Error() string {
	return "dentist not found"
}
