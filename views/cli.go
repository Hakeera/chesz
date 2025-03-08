package views

import (
	"chesz/models"
	"fmt"
)

// Exibir o tabuleiro corretamente de baixo para cima
func PrintBoard(board models.Board) {
    fmt.Println("  a b c d e f g h")
    for i := 7; i >= 0; i-- { // Percorre de baixo para cima
        fmt.Print(i+1, " ") // Exibe o número da linha corretamente
        for _, piece := range board[i] {
            if piece == nil {
                fmt.Print(". ") // Casa vazia
            } else {
                fmt.Print(piece.Type, " ")
            }
        }
        fmt.Println(i + 1)
    }
    fmt.Println("  a b c d e f g h")
}
