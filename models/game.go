package models

import (
	"chesz/views"
	"fmt"
)

type Game struct {
	Board    Board
	Turn     string // "White" ou "Black"
	GameOver bool
}

func NewGame() *Game {
	return &Game{
		Board:    NewBoard(),
		Turn:     "White",
		GameOver: false,
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

func (g *Game) Play() {
	for !g.GameOver {
		views.PrintBoard(g.GetPrintableBoard()) // Usa GetPrintableBoard()

		fmt.Printf("\nTurno: %s\n", g.Turn)

		// Verifica xeque-mate ANTES do turno
		if g.Board.IsCheckmate(g.Turn) {
			fmt.Printf("\nXEQUE-MATE! %s venceu!\n", g.Opponent())
			g.GameOver = true
			return
		}

		fromRow, fromCol, toRow, toCol, err := views.GetMove()

		if err != nil {
			fmt.Println("\nErro:", err)
			continue
		}

		if g.Board.MovePiece(fromRow, fromCol, toRow, toCol, g.Turn) {
			g.SwitchTurn()
		} else {
			views.PrintMessage("Movimento inválido! Tente novamente.")
		}
	}
}

// GetPrintableBoard convert Board to [][]string
func (g *Game) GetPrintableBoard() [][]string {
	printable := make([][]string, 8)
	for i := range printable {
		printable[i] = make([]string, 8)
		for j := range printable[i] {
			if g.Board[i][j] == nil {
				printable[i][j] = "." // Casa vazia
			} else {
				printable[i][j] = g.Board[i][j].Type // Exibir tipo da peça
			}
		}
	}
	return printable
}
