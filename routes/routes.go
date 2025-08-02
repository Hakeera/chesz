// Package routes
package routes

import (
	controller "chesz/controler"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes (e *echo.Echo) {
	e.GET("/", controller.HomeHandler)  
	e.GET("/start", controller.StartGame)  
	e.GET("/move", controller.StartGame)  
}
