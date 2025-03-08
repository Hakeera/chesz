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
    //board.MovePiece(1, 0, 2, 0) 

    fmt.Println("\nMovendo Cavalo") 
    board.MovePiece(0, 1, 2, 2) 

    
    fmt.Println("\nMovendo Cavalo") 
    board.MovePiece(2, 2, 4, 1) 


    fmt.Println("\nMovendo Cavalo") 
    board.MovePiece(4, 1, 5, 3) 

    fmt.Println("\nMovendo Cavalo") 
    board.MovePiece(5, 3, 6, 4) 

    views.PrintBoard(board)
}
