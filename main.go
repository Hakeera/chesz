package main

import (
	"chesz/models"
	"chesz/views"
	"fmt"
)

func main() {
    board := models.NewBoard()
    views.PrintBoard(board)

    //fmt.Println("\nMovendo Peão...")
    //board.MovePiece(1, 1, 2, 1) 
    

    fmt.Println("\nMovendo Cavalo...")
    board.MovePiece(0, 1, 2, 2)
    board.MovePiece(2, 2, 4, 1)
    board.MovePiece(4, 1, 5, 3) 
    fmt.Println("\nXeque!")
    board.MovePiece(5, 3, 7, 4) //Mate!

    //fmt.Println("\nMovendo bispo")
    //board.MovePiece(0, 2, 2, 0)


    //fmt.Println("\nMovendo bispo para captura")
    //board.MovePiece(2, 0, 6, 4)


    views.PrintBoard(board)
    fmt.Println("\nMate!!")
}
