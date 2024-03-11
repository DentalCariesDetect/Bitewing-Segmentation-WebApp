package errors

type SegmentationError struct {
	ServerInternalError       error
	SegmentationNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type ResultNotFoundError struct{}

func (e *ResultNotFoundError) Error() string {
	return "Result not found"
}

type ToothNotFoundError struct{}

func (e *ToothNotFoundError) Error() string {
	return "Tooth not found"
}
