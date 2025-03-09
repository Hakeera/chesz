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

// PrintBoard recebe uma matriz genérica e exibe o tabuleiro
func PrintBoard(board [][]string) {
	fmt.Println("  a b c d e f g h")
	for i, row := range board {
		fmt.Printf("%d ", 8-i)
		for _, cell := range row {
			fmt.Print(cell + " ")
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("  a b c d e f g h")
}

// PrintMessage exibe mensagens no terminal
func PrintMessage(msg string) {
	fmt.Println(msg)
}


// Convert Board to [][]string
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

// Exibir o tabuleiro 
func (g *Game) PrintGameState() {
	views.PrintBoard(g.GetPrintableBoard()) // Chama print.go sem importar models!
}


