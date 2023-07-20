package inject

import (
	echoContext "github.com/daan0220/fib-api/context"
	"github.com/daan0220/fib-api/presentation/fib"

	"github.com/labstack/echo/v4"
)

func FibonacciController() echo.HandlerFunc {
	return newHandlerFunc(fib.Get)

}

type callFunc func(ctx echoContext.Context) error

func newHandlerFunc(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := echoContext.Context{Context: c}
		return h(ctx)
	}
}
