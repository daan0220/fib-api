package common

import (
	echoContext "github.com/daan0220/fib-api/context"
)

func OKWithJsonItem(c echoContext.Context, StatusCode int, item interface{}) error {
	return c.JSON(
		StatusCode,
		item,
	)
}

func ErrorWithJsonItem(c echoContext.Context, StatusCode int, message interface{}) error {
	return c.JSON(
		StatusCode,
		struct {
			Message interface{} `json:"message"`
		}{
			Message: message,
		},
	)
}
