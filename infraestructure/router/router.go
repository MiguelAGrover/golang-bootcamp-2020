package router

import (
	"wizegolangapi/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter this will handle all the routes of the API
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/digimons", func(context echo.Context) error { return c.Digimon.GetDigimons(context) })
	e.POST("/digimons", func(context echo.Context) error { return c.Digimon.CreateDigimon(context) })

	return e
}
