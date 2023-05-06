package fullerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	t.Run("Should return the error text", func(t *testing.T) {
		errorText := "test"

		err := &CustomError{
			StatusCode: 400,
			Err: errors.New(errorText),
		}

		assert.Equal(t, errorText, err.Error())
	})
}

func Test_StatusCode(t *testing.T) {
	t.Run("Should return the error code", func(t *testing.T) {
		errorCode := 400

		err := &CustomError{
			StatusCode: errorCode,
			Err: errors.New("test"),
		}

		assert.Equal(t, errorCode, err.GetStatusCode())		
	})
}