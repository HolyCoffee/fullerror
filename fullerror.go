package fullerror

type CustomError struct {
	StatusCode int
	Err error
}

func (e *CustomError) Error() string {
	return e.Err.Error()
}

func (e *CustomError) GetStatusCode() int {
	return e.StatusCode
}
