# fullerror

library to make your errors more informative

## usage

```go
import (
  "fmt"
  "errors"

  "github.com/HolyCoffee/fullerror"
)

func MySuperFunc() error {
  if err := AnotherFunc(); err != nil {
    return &fullerror.CustomError{
      StatusCode: 400,
      Err:        fmt.Errorf("[Custom error]: %w", err),
    }
  }
}

func OtherSuperFunc() {
  if err := MySuperFunc(); err != nil {
    var customErr *fullerror.CustomError

    if errors.As(err, &customErr) {
      fmt.Sprintf("Error status code: %d", customErr.GetStatusCode())
      fmt.Sprintf("Error description: %s", customErr.Error())
    }
  }
}
```