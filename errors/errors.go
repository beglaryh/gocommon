package errors

type ServiceError struct {
	message   string
	errorType ErrorType
}

type ErrorType int

const (
	BadRequest = 400
	NotFound   = 404
	Internal   = 500
)

var DefaultInternalError ServiceError = ServiceError{
	errorType: Internal,
	message:   "internal server error",
}

func New(errorType ErrorType, message string) ServiceError {
	return ServiceError{
		errorType: errorType,
		message:   message,
	}
}

func NewInternal(message string) ServiceError {
	return ServiceError{
		errorType: Internal,
		message:   message,
	}
}

func NewBadRequest(message string) ServiceError {
	return ServiceError{
		errorType: BadRequest,
		message:   message,
	}
}

func NewNotFound(message string) ServiceError {
	return ServiceError{
		errorType: NotFound,
		message:   message,
	}
}

func (se ServiceError) ErrorType() ErrorType {
	return se.errorType
}

func (se ServiceError) Error() string {
	return se.message
}

func (se ServiceError) StatusCode() int {
	return int(se.errorType)
}
