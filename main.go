package main

import (
	"github.com/JonissonGomes/microservice-with-echo/handlers"
	"github.com/labstack/echo"
)

func main() {
	echo := echo.New()
	echo.Logger.Fatal(echo.Start(":9000"))

	echo.GET("/product/:data", handlers.GetProduct)
	echo.POST("/product", handlers.AddProduct)
}
