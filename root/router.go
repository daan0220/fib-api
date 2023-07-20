package root

import (
	"github.com/daan0220/fib-api/inject"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) { e.GET("/fib", inject.FibonacciController()) }
