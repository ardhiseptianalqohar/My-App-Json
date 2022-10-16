package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DataResi(c echo.Context) error {
	result, err := models.DataResi()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanController(c echo.Context) error {
	nomor := c.FormValue("nomor")
	pengirim := c.FormValue("pengirim")
	penerima := c.FormValue("penerima")
	alamat_penerima := c.FormValue("alamat_penerima")
	product := c.FormValue("product")
	product_type := c.FormValue("product_type")
	status_barang := c.FormValue("status_barang")
	estimasi := c.FormValue("estimasi")

	result, err := models.SimpanData(nomor, pengirim, penerima, alamat_penerima, product, product_type, status_barang, estimasi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
