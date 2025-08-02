// Package controller
package controller

import (
	"chesz/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeHandler GET /
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "base", "AQUI")
}

// StartGame GET /start
func StartGame(c echo.Context) error {
	game := models.NewGame()

	// Aqui renderiza o tabuleiro usando GetPrintableBoard
	return c.Render(http.StatusOK, "base", map[string]any{
		"board": game.GetPrintableBoard(),
	})
}

// MovePiece POST /move
func MovePiece(c echo.Context) error {
	// Em breve: lógica para movimentar peças
	return c.String(http.StatusOK, "Movimento ainda não implementado.")
}
