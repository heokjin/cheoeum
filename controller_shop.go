package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const (
	CatagoryTypeShoes string = "shoes"
	CatagoryTypePants string = "pants"
)

func Shop(c echo.Context) error {
	catagoryType := c.Param("catagoryType")
	fmt.Printf("Shop CatagoryId is %s\n", catagoryType)

	switch catagoryType {
	case CatagoryTypeShoes:
		fmt.Println("A")
	case CatagoryTypePants:
		fmt.Println("B")
	default:
		fmt.Println("Z")
	}

	return c.Render(http.StatusOK, "shop.html", "World")
}

func getShoesList() {

}
