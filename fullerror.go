package fullerror

type CustomError struct {
	statusCode int
	err error
}

func (e *CustomError) Error() string {
	return e.err.Error()
}

func (e *CustomError) StatusCode() int {
	return e.statusCode
}
