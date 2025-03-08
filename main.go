package main

import (
	"chesz/models"
	"chesz/views"
	"fmt"
)

func main() {
    board := models.NewBoard()
    views.PrintBoard(board)

    fmt.Println("\nMovendo Peão Brancas")
    board.MovePiece(1, 4, 3, 4)

    fmt.Println("\nMovendo Peão Pretas")
    board.MovePiece(6, 4, 5, 4)

    fmt.Println("\nMovendo Dama Pretas")
    board.MovePiece(7, 3, 4, 6)
    board.MovePiece(4, 6, 3, 6)


    fmt.Println("\nMovendo Rei Brancas")
    board.MovePiece(1, 5, 2, 5) 
    board.MovePiece(0, 4, 1, 4)
	

    views.PrintBoard(board)
}
