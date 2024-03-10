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

type SegmentationNotFoundError struct{}

func (e *SegmentationNotFoundError) Error() string {
	return "Segmentation not found"
}
