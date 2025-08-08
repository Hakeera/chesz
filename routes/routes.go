// Package routes
package routes

import (
	controller "chesz/controler"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes (e *echo.Echo) {
	e.GET("/", controller.HomeHandler)  
	e.GET("/play", controller.StartGame)  
	e.POST("/move", controller.ClientMove)  
}
