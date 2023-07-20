package main

import (
	"github.com/daan0220/fib-api/root"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	root.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
