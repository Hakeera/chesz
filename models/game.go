package models

import (
	"fmt"
	"strings"
)

type Game struct {
	Board    Board
	Turn     string // "White" ou "Black"
	GameOver bool
	MoveChan  chan MoveCommand
}

type MoveCommand struct {
	FromRow, FromCol, ToRow, ToCol int
	ReplyCh chan MoveResult
}

type MoveRequest struct {
	FromRow int
	FromCol int
	ToRow   int
	ToCol   int
	ReplyCh chan MoveResult
}

type MoveResult struct {
	Success   bool
	GameOver  bool
	Message   string
}

var CurrentGame *Game

func NewGame() *Game {
	return &Game{
		Board:    NewBoard(),
		Turn:     "White",
		GameOver: false,
		MoveChan: make(chan MoveCommand), // Cria o Chanel do jogo
	}
}

func (g *Game) SwitchTurn() {
	if g.Turn == "White" {
		g.Turn = "Black"
	} else {
		g.Turn = "White"
	}
}

func (g *Game) Opponent() string {
	if g.Turn == "White" {
		return "Black"
	}
	return "White"
}

func (g *Game) PlayLoop() {
	//Creates Channel to recive requisitions from client
	for !g.GameOver {
		move := <-g.MoveChan

		// Verifica xeque-mate antes do movimento
		if g.Board.IsCheckmate(g.Turn) {
			g.GameOver = true
			move.ReplyCh <- MoveResult{
				Success:  false,
				GameOver: true,
				Message:  fmt.Sprintf("Xeque-mate! %s venceu.", g.Opponent()),
			}
			continue
		}

		ok, _ := g.Board.MovePiece(move.FromRow, move.FromCol, move.ToRow, move.ToCol, g.Turn)
		if ok { 
			g.SwitchTurn()
			move.ReplyCh <- MoveResult{Success: true}
		} else {
			move.ReplyCh <- MoveResult{
				Success: false,
				Message: "Movimento inválido",
			}
		}
	}
}

// GetPrintableBoard convert Board to [][]string
// Peças brancas em MAIÚSCULAS, peças pretas em minúsculas
func (g *Game) GetPrintableBoard() [][]string {
	printable := make([][]string, 8)
	for i := range printable {
		printable[i] = make([]string, 8)
		for j := range printable[i] {
			if g.Board[i][j] == nil {
				printable[i][j] = "." // Casa vazia
			} else {
				piece := g.Board[i][j]
				if piece.Color == "white" || piece.Color == "White" || piece.Color == "WHITE" {
					// Peças brancas em MAIÚSCULAS
					printable[i][j] = strings.ToUpper(piece.Type)
				} else {
					// Peças pretas em minúsculas
					printable[i][j] = strings.ToLower(piece.Type)
				}
			}
		}
	}
	return printable
}

