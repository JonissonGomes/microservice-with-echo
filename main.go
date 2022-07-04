package main

import (
	"github.com/labstack/echo"
)

func main() {
	echo := echo.New()
	echo.Logger.Fatal(echo.Start(":9000"))
}
