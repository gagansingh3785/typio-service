package errors

type ServiceError interface {
	error
	StatusCode() int
	Message() string
	Code() string
}

type serviceError struct {
	err        error
	code       string
	statusCode int
	message    string
}

func (e *serviceError) Error() string {
	return e.err.Error()
}

func (e *serviceError) StatusCode() int {
	return e.statusCode
}

func (e *serviceError) Message() string {
	if e.message == "" {
		return e.err.Error()
	}
	return e.message
}

func (e *serviceError) Code() string {
	return e.code
}

func (e *serviceError) Is4xxError() bool {
	return e.statusCode >= 400 && e.statusCode < 500
}

func NewServiceError(err error, code string, statusCode int, message string) ServiceError {
	return &serviceError{
		err:        err,
		code:       code,
		statusCode: statusCode,
		message:    message,
	}
}
