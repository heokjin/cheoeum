package main

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Index struct {
	SingleCatagory []IndexSingleCatagory
}

type IndexSingleCatagory struct {
	Href         string
	ImgSrc       string
	Price        string
	CatagoryName string
	CatagoryType string
}

func Hello(c echo.Context) error {
	index := Index{}
	singleCatagory := make([]IndexSingleCatagory, 8)
	shoes := IndexSingleCatagory{}
	shoes.CatagoryName = "귀여운 신발"
	shoes.Href = "/shop"
	shoes.ImgSrc = "img/ssam-bg-img/1.png"
	shoes.Price = "From $180"
	shoes.CatagoryType = CatagoryTypeShoes

	shoes2 := IndexSingleCatagory{}
	shoes2.CatagoryName = "aaaaaa"
	shoes2.Href = "/shop"
	shoes2.ImgSrc = "img/ssam-bg-img/2.png"
	shoes2.Price = "From $80"

	shoes3 := IndexSingleCatagory{}
	shoes3.CatagoryName = "aaaaaa"
	shoes3.Href = "/shop"
	shoes3.ImgSrc = "img/ssam-bg-img/3.png"
	shoes3.Price = "From $80"

	shoes4 := IndexSingleCatagory{}
	shoes4.CatagoryName = "aaaaaa"
	shoes4.Href = "/shop"
	shoes4.ImgSrc = "img/ssam-bg-img/4.png"
	shoes4.Price = "From $80"

	shoes5 := IndexSingleCatagory{}
	shoes5.CatagoryName = "aaaaaa"
	shoes5.Href = "/shop"
	shoes5.ImgSrc = "img/ssam-bg-img/5.png"
	shoes5.Price = "From $80"

	shoes6 := IndexSingleCatagory{}
	shoes6.CatagoryName = "aaaaaa"
	shoes6.Href = "/shop"
	shoes6.ImgSrc = "img/ssam-bg-img/6.png"
	shoes6.Price = "From $80"

	shoes7 := IndexSingleCatagory{}
	shoes7.CatagoryName = "aaaaaa"
	shoes7.Href = "/shop"
	shoes7.ImgSrc = "img/ssam-bg-img/7.png"
	shoes7.Price = "From $80"

	shoes8 := IndexSingleCatagory{}
	shoes8.CatagoryName = "aaaaaa"
	shoes8.Href = "/shop"
	shoes8.ImgSrc = "img/ssam-bg-img/8.png"
	shoes8.Price = "From $80"

	singleCatagory[0] = shoes
	singleCatagory[1] = shoes2
	singleCatagory[2] = shoes3
	singleCatagory[3] = shoes4
	singleCatagory[4] = shoes5
	singleCatagory[5] = shoes6
	singleCatagory[6] = shoes7
	singleCatagory[7] = shoes8
	index.SingleCatagory = singleCatagory

	return c.Render(http.StatusOK, "index.html", index)
}
