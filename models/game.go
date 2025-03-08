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

func (g *Game) Play() {
	for !g.GameOver {
		views.PrintBoard(g.Board)
		fmt.Printf("\nTurno: %s\n", g.Turn)

		fromRow, fromCol, toRow, toCol := views.GetMove()
		if g.Board.MovePiece(fromRow, fromCol, toRow, toCol) {
			g.SwitchTurn()
		} else {
			views.PrintMessage("Movimento inválido! Tente novamente.")
		}
	}
}

