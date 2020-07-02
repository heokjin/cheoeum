package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ProductDetails(c echo.Context) error {
	return c.Render(http.StatusOK, "product-details.html", "World")
}

func Cart(c echo.Context) error {
	return c.Render(http.StatusOK, "cart.html", "World")
}

func CheckOut(c echo.Context) error {
	return c.Render(http.StatusOK, "checkout.html", "World")
}
