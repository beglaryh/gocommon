package serviceerror

type ServiceError struct {
	Type    ErrorType
	Message string
}

type ErrorType int

const (
	BadRequest = 400
	NotFound   = 404
	Internal   = 500
)

func New(errotType ErrorType, message string) ServiceError {
	return ServiceError{
		Type:    errotType,
		Message: message,
	}
}

func (se ServiceError) Error() string {
	return se.Message
}

func (se ServiceError) StatusCode() int {
	return int(se.Type)
}
