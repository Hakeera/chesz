// Package controller
package controller

import (
	"chesz/models"
	"fmt"
	"log"
	"net/http"
	"unicode"

	"github.com/labstack/echo/v4"
)

var CurrentGame *models.Game

// HomeHandler GET /
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

// StartGame GET /play
func StartGame(c echo.Context) error {
	game := models.NewGame()
	game.MoveChan = make(chan models.MoveCommand) // Creates Channel
	CurrentGame = game

	go game.PlayLoop()

	return c.Render(http.StatusOK, "base", map[string]any{
		"board": game.GetPrintableBoard(),
	})
}

// ClientMove POST /move
func ClientMove(c echo.Context) error {

	// Movimento recebido do form 
	from := c.FormValue("from")
	to := c.FormValue("to")
	log.Printf("Received move: from=%s, to=%s\n", from, to)

	fromRow, fromCol, err := parsePosition(from)
	if err != nil {
		return c.Render(http.StatusOK, "base", map[string]any{
			"board": models.CurrentGame.GetPrintableBoard(),
			"msg": "Posição de Origem inválida!" ,
	})
	}
	toRow, toCol, err := parsePosition(to)

	if err != nil {
		return c.Render(http.StatusOK, "base", map[string]any{
			"board": models.CurrentGame.GetPrintableBoard(),
			"msg": "Posição de Destino inválida!" ,
	})
	}

	// Cria canal de resposta e envia a jogada
	moveCh := make(chan models.MoveResult)

	if models.CurrentGame == nil {
		return c.String(http.StatusInternalServerError, "Jogo não inicializado")
	}

	models.CurrentGame.MoveChan <- models.MoveCommand{
		FromRow: fromRow,
		FromCol: fromCol,
		ToRow:   toRow,
		ToCol:   toCol,
		ReplyCh: moveCh,
	}

	result := <-moveCh

	return c.Render(http.StatusOK, "base", map[string]any{
		"board": models.CurrentGame.GetPrintableBoard(),
		"msg":   result.Message,
		"turn":   models.CurrentGame.Turn,
	})
}

func parsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("posição inválida")
	}

	colRune := unicode.ToLower(rune(pos[0]))
	rowRune := rune(pos[1])

	col := int(colRune - 'a') // a=0, b=1, ..., h=7
	row := 8 - int(rowRune-'0') // 8->0, 1->7

	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("posição fora do tabuleiro")
	}

	return row, col, nil
}
