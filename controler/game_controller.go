// Package controller
package controller

import (
	"chesz/models"
	"fmt"
	"log"
	"net/http"
	"time"
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
		"turn":  game.Turn,
	})
}

// GetBoard obtém o tabuleiro atualizado - retorna apenas o conteúdo do tabuleiro
func GetBoard(c echo.Context) error {
	if CurrentGame == nil {
		return c.String(http.StatusInternalServerError, "Jogo não inicializado")
	}
	
	return c.Render(http.StatusOK, "board-content", map[string]any{
		"board": CurrentGame.GetPrintableBoard(),
		"turn":  CurrentGame.Turn,
	})
}

// ClientMove POST /move - adaptado para HTMX
func ClientMove(c echo.Context) error {
	// Movimento recebido do form 
	from := c.FormValue("from")
	to := c.FormValue("to")
	log.Printf("Received move: from=%s, to=%s\n", from, to)

	// Verifica se o jogo está inicializado
	if CurrentGame == nil {
		log.Println("Game not initialized")
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": nil,
			"msg":   "Jogo não inicializado. Inicie um novo jogo.",
			"turn":  "white",
		})
	}

	// Parse da posição de origem
	fromRow, fromCol, err := parsePosition(from)
	if err != nil {
		log.Printf("Invalid from position: %s, error: %v", from, err)
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": CurrentGame.GetPrintableBoard(),
			"msg":   "Posição de origem inválida!",
			"turn":  CurrentGame.Turn,
		})
	}

	// Parse da posição de destino
	toRow, toCol, err := parsePosition(to)
	if err != nil {
		log.Printf("Invalid to position: %s, error: %v", to, err)
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": CurrentGame.GetPrintableBoard(),
			"msg":   "Posição de destino inválida!",
			"turn":  CurrentGame.Turn,
		})
	}

	// Cria canal de resposta e envia a jogada
	moveCh := make(chan models.MoveResult, 1) // Buffer para evitar deadlock

	// Envia comando para o game loop
	select {
	case CurrentGame.MoveChan <- models.MoveCommand{
		FromRow: fromRow,
		FromCol: fromCol,
		ToRow:   toRow,
		ToCol:   toCol,
		ReplyCh: moveCh,
	}:
		// Comando enviado com sucesso
	case <-time.After(1 * time.Second):
		log.Println("Timeout sending move command")
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": CurrentGame.GetPrintableBoard(),
			"msg":   "Erro interno: timeout ao enviar movimento",
			"turn":  CurrentGame.Turn,
		})
	}

	// Aguarda resultado com timeout
	select {
	case result := <-moveCh:
		log.Printf("Move result: %s", result.Message)
		log.Printf("Turn: %s", CurrentGame.Turn)
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": CurrentGame.GetPrintableBoard(),
			"msg":   result.Message,
			"turn":  CurrentGame.Turn,
		})
	case <-time.After(5 * time.Second):
		log.Println("Timeout waiting for move result")
		return c.Render(http.StatusOK, "board-content", map[string]any{
			"board": CurrentGame.GetPrintableBoard(),
			"msg":   "Timeout: movimento demorou muito para ser processado",
			"turn":  CurrentGame.Turn,
		})
	}
}

// ResetGame POST /reset - handler para reiniciar o jogo
func ResetGame(c echo.Context) error {
	log.Println("Resetting game...")
	
	// Para o jogo atual se existir
	if CurrentGame != nil {
		// Fecha o canal se ainda estiver aberto
		select {
		case CurrentGame.MoveChan <- models.MoveCommand{}: // Envia comando vazio para finalizar
		default:
		}
	}

	// Cria novo jogo
	game := models.NewGame()
	game.MoveChan = make(chan models.MoveCommand)
	CurrentGame = game
	go game.PlayLoop()

	return c.Render(http.StatusOK, "board-content", map[string]any{
		"board": game.GetPrintableBoard(),
		"msg":   "Jogo reiniciado!",
		"turn":  game.Turn,
	})
}

// GetTurn GET /turn - retorna apenas o indicador de turno
func GetTurn(c echo.Context) error {
	if CurrentGame == nil {
		return c.HTML(http.StatusOK, "Turno: Brancas")
	}
	
	turnText := "Brancas"
	if CurrentGame.Turn == "black" {
		turnText = "Pretas"
	}
	
	return c.HTML(http.StatusOK, fmt.Sprintf("Turno: %s", turnText))
}

// UpdateGameInfo - handler para atualizar informações do jogo (turno + mensagem)
func UpdateGameInfo(c echo.Context) error {
	if CurrentGame == nil {
		return c.JSON(http.StatusOK, map[string]any{
			"turn": "white",
			"turnText": "Brancas",
			"message": "",
		})
	}
	
	turnText := "Brancas"
	if CurrentGame.Turn == "black" {
		turnText = "Pretas"
	}
	
	return c.JSON(http.StatusOK, map[string]any{
		"turn": CurrentGame.Turn,
		"turnText": turnText,
		"message": "", // Você pode adicionar uma mensagem de status aqui se necessário
	})
}

// parsePosition converte notação algébrica (e.g., "e4") para coordenadas da matriz
func parsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("posição deve ter exatamente 2 caracteres, recebido: %s", pos)
	}

	colRune := unicode.ToLower(rune(pos[0]))
	rowRune := rune(pos[1])

	// Verifica se a coluna é válida (a-h)
	if colRune < 'a' || colRune > 'h' {
		return 0, 0, fmt.Errorf("coluna deve ser entre 'a' e 'h', recebido: %c", colRune)
	}

	// Verifica se a linha é válida (1-8)
	if rowRune < '1' || rowRune > '8' {
		return 0, 0, fmt.Errorf("linha deve ser entre '1' e '8', recebido: %c", rowRune)
	}

	col := int(colRune - 'a')     // a=0, b=1, ..., h=7
	row := 8 - int(rowRune-'0')   // 8->0, 7->1, ..., 1->7

	// Double check dos bounds (redundante mas seguro)
	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("posição fora do tabuleiro: row=%d, col=%d", row, col)
	}

	return row, col, nil
}

// Helper function para debug - pode ser removida em produção
func debugPosition(pos string) string {
	row, col, err := parsePosition(pos)
	if err != nil {
		return fmt.Sprintf("%s -> ERROR: %v", pos, err)
	}
	return fmt.Sprintf("%s -> [%d,%d]", pos, row, col)
}
