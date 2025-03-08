package views

import (
	"chesz/models"
	"fmt"
)

// PrintBoard exibe o tabuleiro no terminal
func PrintBoard(board *models.Board) {
	fmt.Println("  a b c d e f g h")
	for i, row := range *board {
		fmt.Printf("%d ", 8-i)
		for _, piece := range row {
			if piece == nil {
				fmt.Print(". ")
			} else {
				fmt.Print(piece.Type, " ")
			}
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("  a b c d e f g h")
}

// PrintMessage exibe mensagens para o jogador
func PrintMessage(msg string) {
	fmt.Println(msg)
}
