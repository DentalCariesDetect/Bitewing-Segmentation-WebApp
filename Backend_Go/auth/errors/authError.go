package errors

type AuthError struct {
	ServerInternalError       error
	UsernameNotFoundError     string
	UsernameAlreadyExistError string
	PasswordIncorrectError    string
	TokenNotAuthorizedError   string
}

type UsernameNotFoundError struct{}

func (e *UsernameNotFoundError) Error() string {
	return "username not found"
}

type UsernameAlreadyExistError struct{}

func (e *UsernameAlreadyExistError) Error() string {
	return "username already exist"
}

type PasswordIncorrectError struct{}

func (e *PasswordIncorrectError) Error() string {
	return "password incorrect"
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type TokenNotAuthorizedError struct{}

func (e *TokenNotAuthorizedError) Error() string {
	return "token not authorized"
}
