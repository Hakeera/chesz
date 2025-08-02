// Package controller
package controller

import (
	"chesz/models"
	"html/template"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/home/home.html",
					"view/layout/base.html",
)

	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar templates: "+err.Error())
	}

	// Executa o template base
	return tmpl.ExecuteTemplate(c.Response(), "home.html", nil)
}

func StartGame(c echo.Context) error  {
tmpl, err := template.ParseFiles("view/game/board.html",
					"view/layout/base.html",
)

	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		return c.String(http.StatusInternalServerError, "Erro ao carregar templates: "+err.Error())
	}

	game := models.NewGame()
	game.Play()

	// Renderiza o tabuleiro inicial
	return tmpl.ExecuteTemplate(c.Response(), "board.html", nil)
}
