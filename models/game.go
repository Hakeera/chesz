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
		views.PrintBoard(g.GetPrintableBoard()) 
		fmt.Printf("\nTurno: %s\n", g.Turn)

		fromRow, fromCol, toRow, toCol, err := views.GetMove()
		if err != nil {
			fmt.Println("\nErro:", err)
			continue
		}
			if g.Board.MovePiece(fromRow, fromCol, toRow, toCol, g.Turn) {
			// Verifica se um rei foi capturado
			if g.IsKingCaptured() {
				g.GameOver = true
				break
			}

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

// Função para exibir o tabuleiro sem criar um ciclo de importação
func (g *Game) PrintGameState() {
	views.PrintBoard(g.GetPrintableBoard()) // Chama print.go sem importar models!
}

//IsKingCaptured verifica se algum dos reis foi capturado.
func (g *Game) IsKingCaptured() bool {
	hasWhiteKing := false
	hasBlackKing := false

	// Percorre o tabuleiro em busca dos reis
	for _, row := range g.Board {

		for _, piece := range row {
			if piece != nil {
				if piece.Type == "K" && piece.Color == "White" {
					hasWhiteKing = true
				} else if piece.Type == "K" && piece.Color == "Black" {
					hasBlackKing = true
				}
			}
		}
	}

	// Se um dos reis foi capturado, o jogo deve terminar
	if !hasWhiteKing {
		fmt.Println("Rei Branco capturado! Pretas vencem!")
		return true
	}
	if !hasBlackKing {
		fmt.Println("Rei Preto capturado! Brancas vencem!")
		return true
	}

	return false
}

