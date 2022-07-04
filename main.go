package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	echo := echo.New()
	echo.Logger.Fatal(echo.Start(":9000"))
}

func GetProduct(c echo.Context) error {
	prodName := c.QueryParam("name")
	prodType := c.QueryParam("type")
	prodCategory := c.QueryParam("category")
	prodSubCategory := c.QueryParam("subcategory")
	prodPrice := c.QueryParam("price")
	dataType := c.Param("datatype")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("O produto %s chegou no formato %s", prodName, dataType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name":        prodName,
			"type":        prodType,
			"category":    prodCategory,
			"subcategory": prodSubCategory,
			"price":       prodPrice,
			"datatype":    dataType,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Verifique se o tipo está no formado String ou JSON"})
	}
}
