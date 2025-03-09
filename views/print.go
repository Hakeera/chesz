package views

import "fmt"

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
