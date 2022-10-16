package routes

import (
	"myapp/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "===SELAMAT DATANG BOS ARDHI===")
	})

	e.GET("/resi", controllers.DataResi)

	e.POST("/resi", controllers.SimpanController)
	return e
}
