package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JonissonGomes/microservice-with-echo/structs"
	"github.com/labstack/echo"
)

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

func AddProduct(c echo.Context) error {

	newProduct := structs.Product{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&newProduct)

	if err != nil {
		log.Fatal("Não foi possível ler requisição")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	return c.String(http.StatusOK, "Produto adicionando!")
}
